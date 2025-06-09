# Alerting

The provided Prometheus Rule templates are a good starting place to monitor your consumer service as well as some portions of the Kafka Connect setup. These prometheues rules can be leveraged through App Interface by defining the variables in the App SRE Observability namespace for your cluster.

// TODO templates should be added to app interface by us, no one will need to take these templates directly, they just need to provide variables and point to template. templates in this repo are for updating/testing if desired to do so outside of App Interface.

## How to Use the Templates

// TODO more explicit App Interface docs with links when not in public repo
**Consumer Service Alerts**:
1. Locate your clusters openshift-customer-monitoring Namespace file in App Interface (openshift-customer-monitoring.CLUSTER_NAME.yml)
2. Define the resource under `openshiftResources`, targeting the template and providing your variable values

**RDS Alerts**:


```yaml
### My Service Name Consumer Alerts
- provider: prometheus-rule
  type: resource-template
  path: /TBD/PATH/TO/OUR/TEMPLATE/LOCATION
  variables:
    service: kessel-inventory
    env: stage
    severity: medium
    app_team: ciam-authz
    dashboard_link: https://test.com
    app_console_link: https://console.com/app
    connector_console_link: https://console.com/connector
    app_sop: https://sop.com
    job_name: kessel-inventory-api
    namespace: kessel-stage
    outbox_topic: outbox.event.kessel.tuples
    connector: kessel-inventory-source-connector
```

3. Create and Submit a PR to App Interface

## How to Test the Template

If you wish to test your alert definition in App Interface using the above process, you can do so following AppSRE's documentation on [testing PrometheusRules](https://gitlab.cee.redhat.com/service/app-interface/-/blob/master/docs/app-sre/prometheus-rules-tests-in-app-interface.md?ref_type=heads#running-prometheus-tests-locally). This same process will also test the PrometheusRuleTests provided with the template.

To test the templates in this repo for proper Jinja formatting and interpolation (to test template updates), a test script is available in the [testing](./testing/) directory

### Test Setup

0. Requires Python 3.12+ and Pip 3

1. Setup Python venv and install requirements:

```shell
cd testing
python3 -m venv venv && source venv/bin/activate && pip3 install -r requirements.txt
```

2. Update the `config.yml` with your variables for testing

### Test Jinja Processing

To test the Jinja templating itself:

1. Test the template: `./render.py ../TEMPLATE_NAME.yml ../config.yml

2. Test the rules file: `./render.py ../TEMPLATE_TESTS_NAME.yml ../config.yml

3. When done, deactivate the venv: `deactivate`
