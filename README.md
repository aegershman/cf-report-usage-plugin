# cf-trueup-plugin

cf-cli plugin showing memory consumption, application instances (AIs), and service instances (SIs) for each org and space you have permission to access.

Reported SIs are for the "pivotal service suite", which as of writing this includes the following:

- RabbitMQ (`p-rabbit`, `p.rabbitmq`)
- Redis (`p.redis`, `p-redis`)
- MySQL (`p.mysql`, `p-mysql`)

## usage

```sh
# report all orgs you have access to
cf trueup-report

# report specific orgs
cf trueup-report -o myorg
```

```txt
Org north-area is consuming 38116 MB of 307200 MB.
  Space development is consuming 1124 MB memory (0%) of org quota.
    3 apps: 2 running 1 stopped
    4 app instances: 3 running, 1 stopped
    0 service instances of type Service Suite
  Space staging is consuming 0 MB memory (0%) of org quota.
    0 apps: 0 running 0 stopped
    0 app instances: 0 running, 0 stopped
    0 service instances of type Service Suite
  Space production is consuming 0 MB memory (0%) of org quota.
    1 apps: 0 running 1 stopped
    1 app instances: 0 running, 1 stopped
    0 service instances of type Service Suite
  Space jigsheth is consuming 0 MB memory (0%) of org quota.
    0 apps: 0 running 0 stopped
    0 app instances: 0 running, 0 stopped
    0 service instances of type Service Suite
Org S1Pdemo14 is consuming 4096 MB of 102400 MB.
  Space development is consuming 1024 MB memory (1%) of org quota.
    2 apps: 1 running 1 stopped
    2 app instances: 1 running, 1 stopped
    1 service instances of type Service Suite
  Space IoT-ConnectedCar-Emulator is consuming 1024 MB memory (1%) of org quota.
    1 apps: 1 running 0 stopped
    1 app instances: 1 running, 0 stopped
    1 service instances of type Service Suite
  Space sandbox is consuming 0 MB memory (0%) of org quota.
    1 apps: 0 running 1 stopped
    2 app instances: 0 running, 2 stopped
    0 service instances of type Service Suite
  Space auto-2 is consuming 0 MB memory (0%) of org quota.
    1 apps: 0 running 1 stopped
    3 app instances: 0 running, 3 stopped
    0 service instances of type Service Suite
  Space scdf-twitter-demo is consuming 0 MB memory (0%) of org quota.
    7 apps: 0 running 7 stopped
    7 app instances: 0 running, 7 stopped
    0 service instances of type Service Suite
  Space scs-demo is consuming 0 MB memory (0%) of org quota.
    0 apps: 0 running 0 stopped
    0 app instances: 0 running, 0 stopped
    0 service instances of type Service Suite
  Space scdf-twitter-demo-s1p-2018 is consuming 0 MB memory (0%) of org quota.
    0 apps: 0 running 0 stopped
    0 app instances: 0 running, 0 stopped
    0 service instances of type Service Suite
You have deployed 16 apps across 2 org(s), with a total of 20 app instances configured. You are currently running 4 apps with 5 app instances and using 2 service instances of type Service Suite.
```
