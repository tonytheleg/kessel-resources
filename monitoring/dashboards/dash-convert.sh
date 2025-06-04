#!/bin/bash

JSON_FILE=$1

DASHBOARD_ID="12345"
TOPIC_NAME="outbox.events.kessel.tuples"
JOB_NAME="kessel-inventory-api"
DATASOURCE_REGEX_STATEMENT='(crcp01ue1-prometheus)|(crcfrp01ugw1-prometheus)|(crcs02ue1-prometheus)|(crcfrs01ugw1-prometheus)'
SERVICE_NAME="Kessel Inventory"
DASHBOARD_UID="abc123"

sed "s/{{DASHBOARD_ID}}/$DASHBOARD_ID/g; s/{{TOPIC_NAME}}/$TOPIC_NAME/g; s/{{JOB_NAME}}/$JOB_NAME/g; s/{{DATASOURCE_REGEX_STATEMENT}}/$DATASOURCE_REGEX_STATEMENT/g; s/{{SERVICE_NAME}}/$SERVICE_NAME/g; s/{{DASHBOARD_UID}}/$DASHBOARD_UID/g" $JSON_FILE > complete.json
