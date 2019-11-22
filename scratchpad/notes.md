# notes

```txt
cf services

name                  service              plan            bound apps                                     last operation     broker                    upgrade available
autoscaler            app-autoscaler       standard                                                       create succeeded   app-autoscaler
config                p-config-server      standard        cfenv-demo                                     create succeeded   p-spring-cloud-services
config-server         p-config-server      standard        gwx-spring-boot-example-basic                  update succeeded   p-spring-cloud-services
identity              p-identity           uaa             cfenv-demo                                     create succeeded   identity-service-broker
metrics-forwarder     metrics-forwarder    unlimited                                                      create succeeded   metrics-forwarder
metrics-forwarder-1   metrics-forwarder    unlimited       cf-sample-app-nodejs, cfenv-demo, nodejs-web   create succeeded   metrics-forwarder
newrelic              user-provided                        gwx-spring-boot-example-basic
registry              p-service-registry   standard                                                       create succeeded   p-spring-cloud-services
rmq-clustered         p.rabbitmq           clustered-3.7                                                  create succeeded   rabbitmq-odb
rmq-standard          p-rabbitmq           standard                                                       create succeeded   p-rabbitmq
s3                    aws-s3               standard                                                       create succeeded   aws-services-broker
scheduler             scheduler-for-pcf    standard        nodejs-web                                     create succeeded   scheduler-for-pcf
scs-reg               p.service-registry   standard                                                       create succeeded   scs-service-broker
scs-wow               p.config-server      standard        nodejs-web                                     update succeeded   scs-service-broker
small-redis           p.redis              cache-small                                                    create succeeded   redis-odb
tiny-sql              p.mysql              db-small        cfenv-demo                                     update succeeded   dedicated-mysql-broker
```

services that should count:

```txt
# data services, x4
rmq-clustered
rmq-standard
small-redis
tiny-sql

# scs services, x5
config
config-server
registry
scs-reg
scs-wow
```

```txt
cf apps

cf-sample-app-nodejs            started           1/1         512M     1G
cfenv-demo                      started           1/1         1G       256M
gwx-asp-net-core-app-basic      started           1/1         1G       1G
gwx-spring-boot-example-basic   started           1/1         1G       512M
gwx-spring-boot-sample-ui       stopped           0/1         2G       512M
hammerdb-test                   stopped           0/1         1G       1G
nodejs-web                      started           2/2         32M      1G
push-test-webhook-switchboard   started           2/2         1G       1G
```

then let's check it against what's reported

```txt
cf trueup-view -o x

Org x is consuming 12864 MB of 20480 MB.
        Space docs is consuming 2048 MB memory (10%) of org quota.
                0 apps: 0 running 0 stopped
                2 app instances: 2 running, 0 stopped
                0 service instances of type Service Suite
        Space probots is consuming 6144 MB memory (30%) of org quota.
                3 apps: 3 running 0 stopped
                4 app instances: 4 running, 0 stopped
                0 service instances of type Service Suite
        Space route53-sync is consuming 0 MB memory (0%) of org quota.
                0 apps: 0 running 0 stopped
                0 app instances: 0 running, 0 stopped
                0 service instances of type Service Suite
        Space scratchpad is consuming 8768 MB memory (42%) of org quota.
                8 apps: 6 running 2 stopped
                13 app instances: 11 running, 2 stopped
                4 service instances of type Service Suite
        Space splunk-firehose is consuming 1024 MB memory (5%) of org quota.
                1 apps: 1 running 0 stopped
                2 app instances: 2 running, 0 stopped
                0 service instances of type Service Suite
        Space sso is consuming 0 MB memory (0%) of org quota.
                0 apps: 0 running 0 stopped
                0 app instances: 0 running, 0 stopped
                0 service instances of type Service Suite
You have deployed 12 apps across 1 org(s), with a total of 21 app instances configured. You are currently running 10 apps with 19 app instances and using 4 service instances of type Service Suite.
```

We should really only care about 'scratchpad'

```txt
8 apps: 6 running 2 stopped
13 app instances: 11 running, 2 stopped
4 service instances of type Service Suite
```

## what we can deduce

The first line, `8 apps: 6 running 2 stopped`, has nothing to do with app instances at all.

The second line, `13 app instances: 11 running, 2 stopped`, is bizzare. There are `10` app_instances possible if you add up the results from `cf apps`. So it's definitely adding up something extra.

Why should it report like that, though? Let's try something different. My boss wants to know:

- _billable_ app instances
- _billable_ service instances

In order to get that data, I need to know:

- AIs that are _actually_ applications
- AIs that aren't _actually_ applications, but are instead _counted_ as AIs (e.g. SCS, etc., although getting a straight answer on what counts as an AI/SI is surprisingly annoying)
- SIs that are _actually_ service instances
- SIs that aren't _actually_ services, but are instead being _counted_ as AIs
- SIs that are composed of _multiple_ other AIs and SIs, e.g. SCDF
- How _user provided services_ should be counted
