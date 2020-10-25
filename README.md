# cf-report-usage-plugin

cf-cli plugin showing memory consumption, application instances (AIs), and service instances (SIs) for each org and space you have permission to access.

## usage

```sh
# report all orgs you have access to
cf report-usage

# report specific orgs
cf report-usage -o voyager
cf report-usage -o voyager -o tenzing [-o orgName...]

# report using specific formats
cf report-usage -o voyager -o tenzing --format table
cf report-usage -o voyager -o tenzing --format string
cf report-usage -o voyager -o tenzing --format json

# or output multiple report formats in the same run by specifying --format multiple times
cf report-usage -o voyager -o tenzing --format table --format json
cf report-usage -o voyager -o tenzing [--format formatName...]
```

`--format table`:

```txt
+---------+-------------+--------------+-----+-------------+------+-----+
|   ORG   |    SPACE    | BILLABLE AIS | AIS | STOPPED AIS | APPS | SCS |
+---------+-------------+--------------+-----+-------------+------+-----+
| voyager | dev         | 11           | 10  | 1           | 12   | 2   |
| voyager | test        | 10           | 10  | 1           | 11   | 1   |
| tenzing | dev         | 8            | 7   | 1           | 8    | 2   |
| tenzing | integration | 1            | 0   | 0           | 4    | 1   |
| tenzing | test        | 8            | 6   | 0           | 8    | 2   |
+---------+-------------+--------------+-----+-------------+------+-----+
|    -    |    TOTAL    |      38      | 33  |      3      |  43  |  8  |
+---------+-------------+--------------+-----+-------------+------+-----+
```

`--format string`:

```txt
org voyager is consuming 36864 MB of 83968 MB
        space dev is consuming 18432 MB memory of org quota
                AIs billable: 11
                AIs canonical: 10 (9 running, 1 stopped)
                SCS instances: 2
        space test is consuming 18432 MB memory of org quota
                AIs billable: 10
                AIs canonical: 10 (9 running, 1 stopped)
                SCS instances: 1
org tenzing is consuming 31128 MB of 83968 MB
        space dev is consuming 15564 MB memory of org quota
                AIs billable: 8
                AIs canonical: 7 (6 running, 1 stopped)
                SCS instances: 2
        space integration is consuming 0 MB memory of org quota
                AIs billable: 1
                AIs canonical: 0 (0 running, 0 stopped)
                SCS instances: 1
        space test is consuming 15564 MB memory of org quota
                AIs billable: 8
                AIs canonical: 6 (6 running, 0 stopped)
                SCS instances: 2
across 2 org(s), you have 38 billable AIs, 33 are canonical AIs (30 running, 3 stopped), 8 are SCS instances
```

`--format json`:

This example is going to be a bit cluttered, so it's recommended using `jq` to parse the output. Then it could be used for parsing and interacting with other systems, like `cf-mgmt`:

```json
{
  "summary_report": {
    "org_reports": [
      {
        "org_quota": {
          "app_instance_limit": 38,
          "app_task_limit": -1,
          "guid": "7cdf9c89-f493-417e-97bd-5422f9681759",
          "instance_memory_limit": -1,
          "memory_limit": 83968,
          "name": "voyager",
          "total_private_domains": -1,
          "total_reserved_route_ports": 0,
          "total_routes": -1,
          "total_service_keys": -1,
          "total_services": -1
        },
        "app_instances_count": 20,
        "apps_count": 23,
        "billable_app_instances_count": 21,
        "billable_services_count": 11,
        "memory_quota": 83968,
        "memory_usage": 36864,
        "name": "voyager",
        "running_app_instances_count": 18,
        "running_apps_count": 18,
        "services_count": 14,
        "services_suite_for_pivotal_platform_count": 7,
        "spring_cloud_services_count": 3,
        "stopped_app_instances_count": 2,
        "stopped_apps_count": 5,
        "space_reports": [
          {
            "app_instances_count": 10,
            "apps_count": 12,
            "billable_app_instances_count": 11,
            "billable_services_count": 6,
            "memory_quota": -1,
            "memory_usage": 18432,
            "name": "dev",
            "running_app_instances_count": 9,
            "running_apps_count": 9,
            "services_count": 8,
            "services_suite_for_pivotal_platform_count": 4,
            "spring_cloud_services_count": 2,
            "stopped_app_instances_count": 1,
            "stopped_apps_count": 3
          },
          {
            "app_instances_count": 10,
            "apps_count": 11,
            "billable_app_instances_count": 10,
            "billable_services_count": 5,
            "memory_quota": -1,
            "memory_usage": 18432,
            "name": "test",
            "running_app_instances_count": 9,
            "running_apps_count": 9,
            "services_count": 6,
            "services_suite_for_pivotal_platform_count": 3,
            "spring_cloud_services_count": 1,
            "stopped_app_instances_count": 1,
            "stopped_apps_count": 2
          }
        ]
      },
      {
        "org_quota": {
          "app_instance_limit": 28,
          "app_task_limit": -1,
          "guid": "b87e4524-43a3-4e65-9593-013fa0fefa42",
          "instance_memory_limit": -1,
          "memory_limit": 83968,
          "name": "tenzing",
          "total_private_domains": -1,
          "total_reserved_route_ports": 5,
          "total_routes": -1,
          "total_service_keys": -1,
          "total_services": -1
        },
        "app_instances_count": 13,
        "apps_count": 20,
        "billable_app_instances_count": 17,
        "billable_services_count": 15,
        "memory_quota": 83968,
        "memory_usage": 31128,
        "name": "tenzing",
        "running_app_instances_count": 12,
        "running_apps_count": 12,
        "services_count": 20,
        "services_suite_for_pivotal_platform_count": 9,
        "spring_cloud_services_count": 5,
        "stopped_app_instances_count": 1,
        "stopped_apps_count": 8,
        "space_reports": [
          {
            "app_instances_count": 7,
            "apps_count": 8,
            "billable_app_instances_count": 8,
            "billable_services_count": 5,
            "memory_quota": -1,
            "memory_usage": 15564,
            "name": "dev",
            "running_app_instances_count": 6,
            "running_apps_count": 6,
            "services_count": 7,
            "services_suite_for_pivotal_platform_count": 3,
            "spring_cloud_services_count": 2,
            "stopped_app_instances_count": 1,
            "stopped_apps_count": 2
          },
          {
            "app_instances_count": 0,
            "apps_count": 4,
            "billable_app_instances_count": 1,
            "billable_services_count": 5,
            "memory_quota": -1,
            "memory_usage": 0,
            "name": "integration",
            "running_app_instances_count": 0,
            "running_apps_count": 0,
            "services_count": 6,
            "services_suite_for_pivotal_platform_count": 3,
            "spring_cloud_services_count": 1,
            "stopped_app_instances_count": 0,
            "stopped_apps_count": 4
          },
          {
            "app_instances_count": 6,
            "apps_count": 8,
            "billable_app_instances_count": 8,
            "billable_services_count": 5,
            "memory_quota": -1,
            "memory_usage": 15564,
            "name": "test",
            "running_app_instances_count": 6,
            "running_apps_count": 6,
            "services_count": 7,
            "services_suite_for_pivotal_platform_count": 3,
            "spring_cloud_services_count": 2,
            "stopped_app_instances_count": 0,
            "stopped_apps_count": 2
          }
        ]
      }
    ],
    "app_instances_count": 33,
    "apps_count": 43,
    "billable_app_instances_count": 38,
    "billable_services_count": 26,
    "memory_quota": 167936,
    "memory_usage": 67992,
    "name": "voyagertenzing",
    "running_app_instances_count": 30,
    "running_apps_count": 30,
    "services_count": 34,
    "services_suite_for_pivotal_platform_count": 16,
    "spring_cloud_services_count": 8,
    "stopped_app_instances_count": 3,
    "stopped_apps_count": 13
  }
}
```

## installation

If you want to try it out, install it directly from [the github releases tab as follows](https://github.com/aegershman/cf-report-usage-plugin/releases):

```sh
# osx 64bit
cf install-plugin -f https://github.com/aegershman/cf-report-usage-plugin/releases/download/3.3.2/cf-report-usage-plugin-darwin

# linux 64bit (32bit and ARM6 also available)
cf install-plugin -f https://github.com/aegershman/cf-report-usage-plugin/releases/download/3.3.2/cf-report-usage-plugin-linux-amd64

# windows 64bit (32bit also available)
cf install-plugin -f https://github.com/aegershman/cf-report-usage-plugin/releases/download/3.3.2/cf-report-usage-plugin-windows-amd64.exe
```

## backwards compatibility

To be honest, I wouldn't describe this plugin as "totally ready" yet. It's not where I want it yet. I will do the best I can to maintain backwards compatibility with the current set of properties that can be rendered by a presenter.

## use in pivotal licensing

This plugin's usefulness for reporting things Pivotal's licensing on AI/SI packs and such is definitely a work in progress. I'd like to make this more dynamic.

Reported SIs are for the "pivotal service suite", which as of writing this includes the following:

- RabbitMQ (`p-rabbit`, `p.rabbitmq`)
- Redis (`p.redis`, `p-redis`)
- MySQL (`p.mysql`, `p-mysql`)

Services part of the "spring cloud config" (SCS) suite, although are "SIs" from the perspective of CF, are treated as AIs from the perspective of billing. The following service instances are _billed_ and currently in this tool _reported_ as running AIs:

- Spring Cloud Config (`p-spring-cloud-config` in 2.x, `p.spring-cloud-config` in 3.x)
- Service Registry (`p-service-registry` in 2.x, `p.service-registry` in 3.x)
- Circuit Breaker (`p-circuit-breaker` in 2.x, non-existant in 3.x)

## background

This plugin shares the same `git` history as the [`usagereport-plugin`](https://github.com/krujos/usagereport-plugin) and [`trueupreport-plugin`](https://github.com/jigsheth57/trueupreport-plugin). It was forked & [I cleaned up the `git` history to rewrite commits purging files over 1MB](https://rtyley.github.io/bfg-repo-cleaner/) to avoid slow `git` operations.
