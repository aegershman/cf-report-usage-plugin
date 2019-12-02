# cf-usage-report-plugin

cf-cli plugin showing memory consumption, application instances (AIs), and service instances (SIs) for each org and space you have permission to access.

## installation

this plugin isn't currently listed on the cf OSS repo plugin exchange, because I wouldn't really describe it as "totally ready" yet. until then, if you want to try it out, install it directly from [the github releases tab as follows](https://github.com/aegershman/cf-usage-report-plugin/releases):

```sh
# osx 64bit
cf install-plugin -f https://github.com/aegershman/cf-usage-report-plugin/releases/download/2.10.2/cf-usage-report-plugin-darwin

# linux 64bit (32bit and ARM6 also available)
cf install-plugin -f https://github.com/aegershman/cf-usage-report-plugin/releases/download/2.10.2/cf-usage-report-plugin-linux-amd64

# windows 64bit (32bit also available)
cf install-plugin -f https://github.com/aegershman/cf-usage-report-plugin/releases/download/2.10.2/cf-usage-report-plugin-windows-amd64.exe
```

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
+---------+-------------+--------------+-----+-------------+------+-----+
|   ORG   |    SPACE    | BILLABLE AIS | AIS | STOPPED AIS | APPS | SCS |
+---------+-------------+--------------+-----+-------------+------+-----+
| voyager | dev         | 20           | 18  | 5           | 16   | 2   |
| voyager | test        | 20           | 19  | 2           | 12   | 1   |
| tenzing | dev         | 2            | 0   | 0           | 8    | 2   |
| tenzing | test        | 2            | 0   | 0           | 9    | 2   |
| tenzing | integration | 1            | 0   | 0           | 4    | 1   |
+---------+-------------+--------------+-----+-------------+------+-----+
|    -    |    TOTAL    |      45      | 37  |      7      |  49  |  8  |
+---------+-------------+--------------+-----+-------------+------+-----+
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
{"summary_report":{"org_reports":[{"app_instances_count":37,"apps_count":28,"billable_app_instances_count":40,"billable_services_count":13,"memory_quota":83968,"memory_usage":60416,"name":"voyager","running_app_instances_count":30,"running_apps_count":21,"services_count":16,"services_suite_for_pivotal_platform_count":7,"spring_cloud_services_count":3,"stopped_app_instances_count":7,"stopped_apps_count":7,"space_reports":[{"app_instances_count":18,"apps_count":16,"billable_app_instances_count":20,"billable_services_count":7,"memory_quota":-1,"memory_usage":25600,"name":"dev","running_app_instances_count":13,"running_apps_count":11,"services_count":9,"services_suite_for_pivotal_platform_count":4,"spring_cloud_services_count":2,"stopped_app_instances_count":5,"stopped_apps_count":5},{"app_instances_count":19,"apps_count":12,"billable_app_instances_count":20,"billable_services_count":6,"memory_quota":-1,"memory_usage":34816,"name":"test","running_app_instances_count":17,"running_apps_count":10,"services_count":7,"services_suite_for_pivotal_platform_count":3,"spring_cloud_services_count":1,"stopped_app_instances_count":2,"stopped_apps_count":2}]},{"app_instances_count":0,"apps_count":21,"billable_app_instances_count":5,"billable_services_count":18,"memory_quota":83968,"memory_usage":0,"name":"tenzing","running_app_instances_count":0,"running_apps_count":0,"services_count":23,"services_suite_for_pivotal_platform_count":9,"spring_cloud_services_count":5,"stopped_app_instances_count":0,"stopped_apps_count":21,"space_reports":[{"app_instances_count":0,"apps_count":8,"billable_app_instances_count":2,"billable_services_count":6,"memory_quota":-1,"memory_usage":0,"name":"dev","running_app_instances_count":0,"running_apps_count":0,"services_count":8,"services_suite_for_pivotal_platform_count":3,"spring_cloud_services_count":2,"stopped_app_instances_count":0,"stopped_apps_count":8},{"app_instances_count":0,"apps_count":9,"billable_app_instances_count":2,"billable_services_count":6,"memory_quota":-1,"memory_usage":0,"name":"test","running_app_instances_count":0,"running_apps_count":0,"services_count":8,"services_suite_for_pivotal_platform_count":3,"spring_cloud_services_count":2,"stopped_app_instances_count":0,"stopped_apps_count":9},{"app_instances_count":0,"apps_count":4,"billable_app_instances_count":1,"billable_services_count":6,"memory_quota":-1,"memory_usage":0,"name":"integration","running_app_instances_count":0,"running_apps_count":0,"services_count":7,"services_suite_for_pivotal_platform_count":3,"spring_cloud_services_count":1,"stopped_app_instances_count":0,"stopped_apps_count":4}]}],"app_instances_count":37,"apps_count":49,"billable_app_instances_count":45,"billable_services_count":31,"memory_quota":167936,"memory_usage":60416,"name":"voyagertenzing","running_app_instances_count":30,"running_apps_count":21,"services_count":39,"services_suite_for_pivotal_platform_count":16,"spring_cloud_services_count":8,"stopped_app_instances_count":7,"stopped_apps_count":28},"format":"json"}
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
