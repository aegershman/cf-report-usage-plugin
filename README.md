# cf-trueup-plugin

This CF CLI Plugin to shows memory consumption and application & service instances (only part of service suite (RabbitMQ, Redis & MySQL)) for each org and space you have permission to access.

## usage

For human readable output:

```txt
➜  trueupreport-plugin git:(master) ✗ cf trueup-report
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

CSV output:

```txt
Creates sqllite db (usagereport.db) in working directory, so you can do offline BI!
➜  trueupreport-plugin git:(master) ✗ cf trueup-report -f csv

➜  trueupreport-plugin git:(master) ✗ cf trueup-report -f csv
Env, ReportDate, OrgName, SpaceName, SpaceMemoryUsed, OrgMemoryQuota, AppsDeployed, AppsRunning, AppInstancesConfigured, AppInstancesRunning, TotalServiceInstancesDeployed, RabbitMQServiceInstanceDeployed, RedisServiceInstanceDeployed, MySQLServiceInstanceDeployed, SpringCloudServiceInstanceDeployed, SpringCloudDataFlowServerInstanceDeployed
api.run.pivotal.io, 2018-12-08, north-area, development, 1124, 307200, 3, 2, 4, 3, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, north-area, staging, 0, 307200, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, north-area, production, 0, 307200, 1, 0, 1, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, north-area, jigsheth, 0, 307200, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, development, 1024, 102400, 2, 1, 2, 1, 1, 1, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, IoT-ConnectedCar-Emulator, 1024, 102400, 1, 1, 1, 1, 1, 0, 1, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, sandbox, 0, 102400, 1, 0, 2, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, auto-2, 0, 102400, 1, 0, 3, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, scdf-twitter-demo, 0, 102400, 7, 0, 7, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, scs-demo, 0, 102400, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
api.run.pivotal.io, 2018-12-08, S1Pdemo14, scdf-twitter-demo-s1p-2018, 0, 102400, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0
```

## install from source

```sh
go get github.com/mattn/go-sqlite3
go get github.com/cloudfoundry/cli
go get github.com/jigsheth57/trueupreport-plugin
cd $GOPATH/src/github.com/jigsheth57/trueupreport-plugin
go build
cf install-plugin trueupreport-plugin
```
