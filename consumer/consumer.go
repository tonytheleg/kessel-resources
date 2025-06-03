package main

import (
	"context"
	"fmt"
	"os"
	"os/signal"
	"syscall"
	"time"

	"github.com/confluentinc/confluent-kafka-go/v2/kafka"
	kessel "github.com/project-kessel/inventory-api/api/kessel/inventory/v1beta1/resources"
	"github.com/project-kessel/inventory-client-go/common"
	"github.com/project-kessel/inventory-client-go/v1beta1"
)

func main() {

	c, err := kafka.NewConsumer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		// "sasl.username":     "inventory-consumer",
		// "sasl.password":     "<REPLACE-ME>",

		// "security.protocol": "sasl_plaintext",
		// "sasl.mechanisms":   "SCRAM-SHA-512",
		"group.id":          "inventory-consumer",
		"auto.offset.reset": "earliest"})

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
			ev, err := c.ReadMessage(100 * time.Millisecond)
			if err == nil {
				fmt.Printf("Consumed event from topic %s: key = %-10s value = %s\n",
					*ev.TopicPartition.Topic, string(ev.Key), string(ev.Value))

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

				// send out request
				_, err = client.RhelHostServiceClient.CreateRhelHost(context.Background(), &request, opts...)
				if err != nil {
					fmt.Println(err)
				}

			} else if !err.(kafka.Error).IsTimeout() {
				fmt.Printf("Consumer error: %v (%v)\n", err, ev)
			}
		}
	}

	c.Close()
}
