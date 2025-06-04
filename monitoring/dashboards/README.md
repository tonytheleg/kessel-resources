# Dashboards
// TODO we'll need to provide docs on custom metrics and stats metrics used and recommended for monitoring

This is not a fully fleshed out implementation, but a sample of what we could do for templating dashboards. The `consumer-dashboard.json` file contains a dashboard definition that pulls all the counsumer related metrics from our Inventory API dashboard and parameterizes some of the values.

The `dash-convert.sh` script is an example of how we could provide a tool for teams to update the variables in the json file using basic cli tools. Its currently hardcoded with values to change, its just meant to be a breadcrumb idea of how we can implement it. It will require follow up work.

Taking the script and json file, they can easily update the values to theirs then dump it into a configmap needed by App Interface. It would be great if we could handle that step for them in the script, but attempts to merge the two have been problematic since it has to be a multiline string and it kept stripping out the leading bracket

The template does not account for the permanent metric names we expect teams to use yet, as we hash that out we can improve upon this idea.
