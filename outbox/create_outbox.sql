-- This SQL script creates an outbox table for Debezium to capture events from
-- id, the primary unique identifier for the event
-- aggregatetype, the type of aggregate that the event belongs to (e.g., "cluster", "user"). This field is used to tell Debezium to route to different Kafka topics.
-- aggregateid, the unique identifier of the aggregate instance (e.g., a grouping identifier). This is important for maintaining correct order in Kafka partitions.
-- payload, the JSON payload containing the event data (typically a full "resource" or API request body)
-- Please review the following documentation for more information on the outbox table:
-- https://debezium.io/documentation/reference/stable/transformations/outbox-event-router.html#basic-outbox-table
CREATE TABLE outbox (
    id UUID NOT NULL,
    aggregatetype VARCHAR(255) NOT NULL,
    aggregateid VARCHAR(255) NOT NULL,
    payload JSONB,
    PRIMARY KEY (id)
);
-- Any additional columns from the outbox table can be added to outbox events either within the payload section or as a message header.
-- One example could be a column eventType which conveys a user-defined value that helps to categorize or organize events.