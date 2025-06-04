# Alerting

From an alerting standpoint we can easily convert our existing prometheus rules and tests to jinja templates. AppSRE has a prescribed method of interpolating and applying these templates to create the prometheus rules in their namespaces, so we have plenty of examples we can follow on converting ours

These could likely live in App Interface vs this resources repo so that they can just be consumed by passing params and not force people to copy the template over
