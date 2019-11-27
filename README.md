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
cf usage-report -o voyager -o tenzing --format table
cf usage-report -o voyager -o tenzing --format string
cf usage-report -o voyager -o tenzing --format json
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
        space dev is consuming 25600 MB memory of org quota
                AIs billable: 20
                AIs canonical: 18 (13 running, 5 stopped)
                SCS instances: 2
        space test is consuming 34816 MB memory of org quota
                AIs billable: 20
                AIs canonical: 19 (17 running, 2 stopped)
                SCS instances: 1
org tenzing is consuming 0 MB of 83968 MB
        space dev is consuming 0 MB memory of org quota
                AIs billable: 2
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 2
        space test is consuming 0 MB memory of org quota
                AIs billable: 2
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 2
        space integration is consuming 0 MB memory of org quota
                AIs billable: 1
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 1
across 2 org(s), you have 45 billable AIs, 37 are canonical AIs (30 running, 7 stopped), 8 are SCS instances
```

`--format json`:

This example is going to be a bit cluttered, so it's recommended using `jq` to parse the output. Then it could be used for parsing and interacting with other systems, like `cf-mgmt`:

```json
{"SummaryReport":{"OrgReports":[{"AppInstancesCount":37,"AppsCount":28,"BillableAppInstancesCount":40,"BillableServicesCount":13,"MemoryQuota":83968,"MemoryUsage":60416,"Name":"voyager","RunningAppInstancesCount":30,"RunningAppsCount":21,"ServicesCount":16,"ServicesSuiteForPivotalPlatformCount":7,"SpringCloudServicesCount":3,"StoppedAppInstancesCount":7,"StoppedAppsCount":7,"SpaceReports":[{"AppInstancesCount":18,"AppsCount":16,"BillableAppInstancesCount":20,"BillableServicesCount":7,"MemoryQuota":-1,"MemoryUsage":25600,"Name":"dev","RunningAppInstancesCount":13,"RunningAppsCount":11,"ServicesCount":9,"ServicesSuiteForPivotalPlatformCount":4,"SpringCloudServicesCount":2,"StoppedAppInstancesCount":5,"StoppedAppsCount":5},{"AppInstancesCount":19,"AppsCount":12,"BillableAppInstancesCount":20,"BillableServicesCount":6,"MemoryQuota":-1,"MemoryUsage":34816,"Name":"test","RunningAppInstancesCount":17,"RunningAppsCount":10,"ServicesCount":7,"ServicesSuiteForPivotalPlatformCount":3,"SpringCloudServicesCount":1,"StoppedAppInstancesCount":2,"StoppedAppsCount":2}]},{"AppInstancesCount":0,"AppsCount":21,"BillableAppInstancesCount":5,"BillableServicesCount":18,"MemoryQuota":83968,"MemoryUsage":0,"Name":"tenzing","RunningAppInstancesCount":0,"RunningAppsCount":0,"ServicesCount":23,"ServicesSuiteForPivotalPlatformCount":9,"SpringCloudServicesCount":5,"StoppedAppInstancesCount":0,"StoppedAppsCount":21,"SpaceReports":[{"AppInstancesCount":0,"AppsCount":8,"BillableAppInstancesCount":2,"BillableServicesCount":6,"MemoryQuota":-1,"MemoryUsage":0,"Name":"dev","RunningAppInstancesCount":0,"RunningAppsCount":0,"ServicesCount":8,"ServicesSuiteForPivotalPlatformCount":3,"SpringCloudServicesCount":2,"StoppedAppInstancesCount":0,"StoppedAppsCount":8},{"AppInstancesCount":0,"AppsCount":9,"BillableAppInstancesCount":2,"BillableServicesCount":6,"MemoryQuota":-1,"MemoryUsage":0,"Name":"test","RunningAppInstancesCount":0,"RunningAppsCount":0,"ServicesCount":8,"ServicesSuiteForPivotalPlatformCount":3,"SpringCloudServicesCount":2,"StoppedAppInstancesCount":0,"StoppedAppsCount":9},{"AppInstancesCount":0,"AppsCount":4,"BillableAppInstancesCount":1,"BillableServicesCount":6,"MemoryQuota":-1,"MemoryUsage":0,"Name":"integration","RunningAppInstancesCount":0,"RunningAppsCount":0,"ServicesCount":7,"ServicesSuiteForPivotalPlatformCount":3,"SpringCloudServicesCount":1,"StoppedAppInstancesCount":0,"StoppedAppsCount":4}]}],"AppInstancesCount":37,"AppsCount":49,"BillableAppInstancesCount":45,"BillableServicesCount":31,"MemoryQuota":167936,"MemoryUsage":60416,"Name":"voyagertenzing","RunningAppInstancesCount":30,"RunningAppsCount":21,"ServicesCount":39,"ServicesSuiteForPivotalPlatformCount":16,"SpringCloudServicesCount":8,"StoppedAppInstancesCount":7,"StoppedAppsCount":28},"Format":"json"}
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
