package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kessel "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	"github.com/project-kessel/inventory-client-go/common"
	"github.com/project-kessel/inventory-client-go/v1beta1"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"group.id":          "generic-consumer",
		"auto.offset.reset": "earliest"})
	// Below required for auth

	// "sasl.username":     "inventory-consumer",
	// "sasl.password":     "<REPLACE-ME>",

	// "security.protocol": "sasl_plaintext",
	// "sasl.mechanisms":   "SCRAM-SHA-512",

	if err != nil {
		fmt.Printf("failed to create consumer: %v", err)
		os.Exit(1)
	}

	topic := "outbox.event.kessel.tuples" //todo:change
	err = c.SubscribeTopics([]string{topic}, nil)
	if err != nil {
		fmt.Printf("failed to subscribe to topic: %v", err)
		os.Exit(1)
	}

	sigchan := make(chan os.Signal, 1)
	signal.Notify(sigchan, syscall.SIGINT, syscall.SIGTERM)

	client, err := v1beta1.New(common.NewConfig(
		common.WithgRPCUrl("localhost:9081"),
		common.WithTLSInsecure(true),
		common.WithAuthEnabled("", "", ""),
	))
	if err != nil {
		fmt.Println(err)
	}

	opts, err := client.GetTokenCallOption()
	if err != nil {
		fmt.Println(err)
	}

	// Process messages
	run := true
	for run {
		select {
		case sig := <-sigchan:
			fmt.Printf("Caught signal %v: terminating\n", sig)
			run = false
		default:
			event := c.Poll(100)
			if event == nil {
				continue
			}

			switch e := event.(type) {
			case *kafka.Message:
				// some sort of header parsing

				// process messages

				// determine what request type to send to inventory.
				request := kessel.CreateRhelHostRequest{RhelHost: &kessel.RhelHost{
					Metadata: &kessel.Metadata{
						ResourceType: "rhel-host",
						WorkspaceId:  "",
					},
					ReporterData: &kessel.ReporterData{
						ReporterType:       kessel.ReporterData_ACM,
						ReporterInstanceId: "service-account-svc-test",
						ConsoleHref:        "www.example.com",
						ApiHref:            "www.example.com",
						LocalResourceId:    "1",
						ReporterVersion:    "0.1",
					},
				}}

				// send out request to inventory api
				_, err = client.RhelHostServiceClient.CreateRhelHost(context.Background(), &request, opts...)
				if err != nil {
					fmt.Println(err)
				}

				fmt.Printf("consumed event from topic %s, partition %d at offset %s\n",
					*e.TopicPartition.Topic, e.TopicPartition.Partition, e.TopicPartition.Offset)
				fmt.Printf("consumed event data: key = %-10s value = %s\n",
					string(e.Key), string(e.Value))

			case kafka.Error:
				if e.IsFatal() {
					run = false
				} else {
					fmt.Printf("recoverable consumer error: %v (%v)\n", e.Code(), e)
					continue
				}
			case *kafka.Stats:
				// collect metrics here
			default:
				fmt.Printf("event type ignored %v", e)
			}
		}
	}
	c.Close()
}
