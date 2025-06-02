# Kafka Connector Resources

This directory contains the OpenShift template and configuration files needed to deploy a KafkaConnector resource using Streams for Apache Kafka. The Connector CR configures the Debezium for Postgresql connector plugin in your Kafka Connect cluster to monitor changes in your database outbox table and publish events to the defined topic.

### Debezium Property Configs

The Kafka Connector CR template provided in this repo, in its default state, is not complete and requires replacing the config section that sets the Kafka config properties for both the Debezium connector and the Outbox event router.

In the `configs` directory are some default base configs that can be used, and are defined separately for each supported Debezium version. We recommend using these configs in combination with the base Debezium template to generate your own deployable template file for consumption in your deployment pipeline or for Ephemeral testing -- See [Using the Templates](#using-the-templates).

The configuration settings are generally the same across multiple versions, but property names have changed in the past, so it is recommended to use the config specific to the version of Debezium you are running.

### Kafka Connector (Debezium Connector)

The KafkaConnector CR template can be used to deploy a Kafka Connector targeting your Kafka Connect cluster and database. The template is designed to allow for multiple service providers to use the same template while avoiding name duplication or conflicts with topics as these Connectors will likely live in the same namespace and target the same Kafka infrastructure.

As mentioned earlier, the template in its default state is not complete, and requires replacing the config section that sets the Kafka config properties for both the Debezium connector and the Outbox event router.

#### Using the Templates

> [!NOTE]
> This process leverages [yq](https://github.com/mikefarah/yq?tab=readme-ov-file#install) to create the template, you will need to install it

To generate your deployment template:
1. Determine the Debezium version you are using (must be one of the supported versions in `configs` directory)
2. Create your finished OpenShift Template:

```shell
yq eval-all 'select(fileIndex==0).objects[0].spec.config = select(fileIndex==1) | select(fileIndex==0)' \
    kafka-connector/debezium-connector-base.yml \
    kafka-connector/configs/debezium-VERSION.yml > deploy-debezium.yml
```

This produces a complete OpenShift template that can be used locally with the `oc process` command, or added to your existing deployment template by copying the contents of the template to your existing template, and adding the parameters to your existing parameter section in your deploy template.

To use the generated template directly:
```shell
oc process --local -f deploy-debezium.yml \
    -p SERVICE_NAME=<Name of Service Providers Service> \
    -p KAFKA_CONNECT_INSTANCE=<Name of your Kafka Connect Instance> \
    -p DB_SECRET_NAME=<Name of the K8s secret containing DB info (created by App Interface) > | oc apply -f -
```

// TODO -- docs on how to create topics in Platform MQ and what those topics should be called

// TODO -- docs on heartbeat, what topics are needed, why its needed, talk about the WAL growth problem

// TODO -- docs should link to connector docs? outbox docs? other docs that are useful like the KubernetesSecretConfigProvider doc?
