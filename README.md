# cf-trueup-plugin

cf-cli plugin showing memory consumption, application instances (AIs), and service instances (SIs) for each org and space you have permission to access.

## usage

```sh
# report all orgs you have access to
cf trueup-view

# report specific orgs
cf trueup-view -o myorg
cf trueup-view -o firstorg -o secondorg [-o orgName...]

# report using different formats ("string" is default)
cf trueup-view -o firstorg -o secondorg --format string
cf trueup-view -o firstorg -o secondorg --format table
```

Example of `string` format:

```txt
Org firstorg is consuming 12864 MB of 20480 MB.
        Space docs is consuming 0 MB memory (0%) of org quota.
                0 app instances: 0 running 0 stopped
                4 billable app instances (includes AIs and billable SIs, like SCS)
                0 unique app_guids: 0 running 0 stopped
                4 service instances total
                4 service instances of type SCS (config-server, eureka, etc.)
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space probots is consuming 6144 MB memory (30%) of org quota.
                4 app instances: 4 running 0 stopped
                4 billable app instances (includes AIs and billable SIs, like SCS)
                3 unique app_guids: 3 running 0 stopped
                0 service instances total
                0 service instances of type SCS (config-server, eureka, etc.)
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space route53-sync is consuming 0 MB memory (0%) of org quota.
                0 app instances: 0 running 0 stopped
                0 billable app instances (includes AIs and billable SIs, like SCS)
                0 unique app_guids: 0 running 0 stopped
                0 service instances total
                0 service instances of type SCS (config-server, eureka, etc.)
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space scratchpad is consuming 5696 MB memory (27%) of org quota.
                10 app instances: 8 running 2 stopped
                15 billable app instances (includes AIs and billable SIs, like SCS)
                8 unique app_guids: 6 running 2 stopped
                17 service instances total
                5 service instances of type SCS (config-server, eureka, etc.)
                4 service instances of type Service Suite (mysql, redis, rmq)
        Space splunk-firehose is consuming 1024 MB memory (5%) of org quota.
                2 app instances: 2 running 0 stopped
                2 billable app instances (includes AIs and billable SIs, like SCS)
                1 unique app_guids: 1 running 0 stopped
                0 service instances total
                0 service instances of type SCS (config-server, eureka, etc.)
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space sso is consuming 0 MB memory (0%) of org quota.
                0 app instances: 0 running 0 stopped
                0 billable app instances (includes AIs and billable SIs, like SCS)
                0 unique app_guids: 0 running 0 stopped
                1 service instances total
                0 service instances of type SCS (config-server, eureka, etc.)
                0 service instances of type Service Suite (mysql, redis, rmq)
Org secondorg is consuming 30720 MB of 65536 MB.
        Space dev is consuming 30720 MB memory (46%) of org quota.
                30 app instances: 30 running 0 stopped
                31 billable app instances (includes AIs and billable SIs, like SCS)
                30 unique app_guids: 30 running 0 stopped
                3 service instances total
                1 service instances of type SCS (config-server, eureka, etc.)
                2 service instances of type Service Suite (mysql, redis, rmq)
Across 2 org(s), you have 56 billable AIs, 46 are canonical AIs (44 running, 2 stopped), 10 are SCS instances
```

And the `table` format (wip):

```txt
+-------------+-----------------+--------------+--------------+
|     ORG     |      SPACE      | BILLABLE AIS | BILLABLE SIS |
+-------------+-----------------+--------------+--------------+
| firstorg    | docs            | 4            | 0            |
| firstorg    | probots         | 4            | 0            |
| firstorg    | route53-sync    | 0            | 0            |
| firstorg    | scratchpad      | 15           | 12           |
| firstorg    | splunk-firehose | 2            | 0            |
| firstorg    | sso             | 0            | 1            |
| secondorg   | dev             | 31           | 2            |
+-------------+-----------------+--------------+--------------+
|      -      |      TOTAL      |      56      |      15      |
+-------------+-----------------+--------------+--------------+
```

## use in pivotal licensing

This plugin's usefulness for reporting things Pivotal's licensing on AI/SI packs and such is definitely a work in progress

Reported SIs are for the "pivotal service suite", which as of writing this includes the following:

- RabbitMQ (`p-rabbit`, `p.rabbitmq`)
- Redis (`p.redis`, `p-redis`)
- MySQL (`p.mysql`, `p-mysql`)

Services part of the "spring cloud config" (SCS) suite, although are "SIs" from the perspective of CF, are treated as AIs from the perspective of billing. The following service instances are _billed_ and currently in this tool _reported_ as running AIs:

- Spring Cloud Config (`p-spring-cloud-config` in 2.x, `p.spring-cloud-config` in 3.x)
- Service Registry (`p-service-registry` in 2.x, `p.service-registry` in 3.x)
- Circuit Breaker (`p-circuit-breaker` in 2.x, non-existant in 3.x)
