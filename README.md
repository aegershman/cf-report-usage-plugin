# cf-trueup-plugin

(some of the reporting data in this plugin has been wrong; fixing it to be accurate and extensible is a work in progress)

cf-cli plugin showing memory consumption, application instances (AIs), and service instances (SIs) for each org and space you have permission to access.

Reported SIs are for the "pivotal service suite", which as of writing this includes the following:

- RabbitMQ (`p-rabbit`, `p.rabbitmq`)
- Redis (`p.redis`, `p-redis`)
- MySQL (`p.mysql`, `p-mysql`)

Services part of the "spring cloud config" (SCS) suite, although are "SIs" from the perspective of CF, are treated as AIs from the perspective of billing. The following service instances are _billed_ and currently in this tool _reported_ as running AIs:

- Spring Cloud Config (`p-spring-cloud-config` in 2.x, `p.spring-cloud-config` in 3.x)
- Service Registry (`p-service-registry` in 2.x, `p.service-registry` in 3.x)
- Circuit Breaker (`p-circuit-breaker` in 2.x, non-existant in 3.x)

## usage

```sh
# report all orgs you have access to
cf trueup-view

# report specific orgs
cf trueup-view -o myorg
cf trueup-view -o firstorg -o secondorg [-o orgName...]
```

```txt
Org myorg is consuming 12864 MB of 20480 MB.
        Space docs is consuming 4096 MB memory (20%) of org quota.
                0 canonical app instances
                4 billable app instances: 4 running, 0 stopped
                0 unique app_guids: 0 running 0 stopped
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space probots is consuming 6144 MB memory (30%) of org quota.
                4 canonical app instances
                4 billable app instances: 4 running, 0 stopped
                3 unique app_guids: 3 running 0 stopped
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space route53-sync is consuming 0 MB memory (0%) of org quota.
                0 canonical app instances
                0 billable app instances: 0 running, 0 stopped
                0 unique app_guids: 0 running 0 stopped
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space scratchpad is consuming 10816 MB memory (52%) of org quota.
                10 canonical app instances
                15 billable app instances: 13 running, 2 stopped
                8 unique app_guids: 6 running 2 stopped
                4 service instances of type Service Suite (mysql, redis, rmq)
        Space splunk-firehose is consuming 1024 MB memory (5%) of org quota.
                2 canonical app instances
                2 billable app instances: 2 running, 0 stopped
                1 unique app_guids: 1 running 0 stopped
                0 service instances of type Service Suite (mysql, redis, rmq)
        Space sso is consuming 0 MB memory (0%) of org quota.
                0 canonical app instances
                0 billable app instances: 0 running, 0 stopped
                0 unique app_guids: 0 running 0 stopped
                0 service instances of type Service Suite (mysql, redis, rmq)
[WARNING: THIS REPORT SUMMARY IS MISLEADING AND INCORRECT. IT WILL BE FIXED SOON.] You have deployed 12 apps across 1 org(s), with a total of 25 app instances configured. You are currently running 10 apps with 23 app instances and using 4 service instances of type Service Suite.
```
