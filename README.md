# cf-usage-report-plugin

cf-cli plugin showing memory consumption, application instances (AIs), and service instances (SIs) for each org and space you have permission to access.

## usage

```sh
# report all orgs you have access to
cf usage-report

# report specific orgs
cf usage-report -o voyager
cf usage-report -o voyager -o tenzing [-o orgName...]

# report using different formats
cf usage-report -o voyager -o tenzing --format string
cf usage-report -o voyager -o tenzing --format table
```

`--format table`:

```txt
+---------+-------------+--------------+-----+-------------+-----+
|   ORG   |    SPACE    | BILLABLE AIS | AIS | STOPPED AIS | SCS |
+---------+-------------+--------------+-----+-------------+-----+
| voyager | dev         | 20           | 18  | 5           | 2   |
| voyager | test        | 20           | 19  | 2           | 1   |
| tenzing | dev         | 2            | 0   | 0           | 2   |
| tenzing | test        | 2            | 0   | 0           | 2   |
| tenzing | integration | 1            | 0   | 0           | 1   |
+---------+-------------+--------------+-----+-------------+-----+
|    -    |    TOTAL    |      45      | 37  |      7      |  8  |
+---------+-------------+--------------+-----+-------------+-----+
```

`--format string`:

```txt
org voyager is consuming 60416 MB of 83968 MB
        space dev is consuming 25600 MB memory (30%) of org quota
                AIs billable: 20
                AIs canonical: 18 (13 running, 5 stopped)
                SCS instances: 2
        space test is consuming 34816 MB memory (41%) of org quota
                AIs billable: 20
                AIs canonical: 19 (17 running, 2 stopped)
                SCS instances: 1
org tenzing is consuming 0 MB of 83968 MB
        space dev is consuming 0 MB memory (0%) of org quota
                AIs billable: 2
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 2
        space test is consuming 0 MB memory (0%) of org quota
                AIs billable: 2
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 2
        space integration is consuming 0 MB memory (0%) of org quota
                AIs billable: 1
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 1
across 2 org(s), you have 45 billable AIs, 37 are canonical AIs (30 running, 7 stopped), 8 are SCS instances
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
