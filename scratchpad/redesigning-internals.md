# redesigning-internals

before we get too far into it, here's result of report-usage output on 2.6.0

```txt
Org x is consuming 12864 MB of 20480 MB.
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
