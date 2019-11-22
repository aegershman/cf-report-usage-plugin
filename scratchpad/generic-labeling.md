# generic-labeling

instead of putting the rules directly in the code, can they be externalized?

e.g., a way to define...

- "any `user-provided-services` should be counted as an AI"
- "any service which comes from `<some-broker>` should be reported as an AI"
- "anything which has this metadata information should be treated as xyzabc"
- orgs to be filtered out of _all_ searches, e.g. like how in models/report there's this line:

```go
go orgStats.Spaces.Stats(chSpaceStats, orgStats.Name == "p-spring-cloud-services")
```

overthinking example:

```yml
rules:
- id: someruleid
  category: {services, apps}
  match: # all optional matching
  - service: servicename
    plan: serviceplan
    service_instance: serviceinstanceregex
    broker: servicebrokername
  report_as:
    category: {services, apps, none}
    # would still need some way to have SCDF report as three different apps
```
