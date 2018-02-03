# TrueupReport Plugin
This CF CLI Plugin to shows memory consumption and application instances for each org and space you have permission to access.

[![wercker status](https://app.wercker.com/status/8881b5530809e3636080d2df6433aada/s/master "wercker status")](https://app.wercker.com/project/bykey/8881b5530809e3636080d2df6433aada)


#Usage

For human readable output:

```
➜  trueupreport-plugin git:(master) ✗ cf trueup-report
Gathering usage information
Org platform-eng is consuming 53400 MB of 204800 MB.
	Space CFbook is consuming 128 MB memory (0%) of org quota.
		1 apps: 1 running 0 stopped
		1 app instances: 1 running, 0 stopped
		0 service instances of type Service Suite
Org krujos is consuming 512 MB of 10240 MB.
	Space development is consuming 0 MB memory (0%) of org quota.
		4 apps: 0 running 4 stopped
		4 app instances: 0 running, 4 stopped
		1 service instances of type Service Suite
	Space production is consuming 512 MB memory (5%) of org quota.
		1 apps: 1 running 0 stopped
		2 app instances: 2 running, 0 stopped
		0 service instances of type Service Suite
Org pcfp is consuming 7296 MB of 102400 MB.
	Space development is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space docs-staging is consuming 512 MB memory (0%) of org quota.
		2 apps: 1 running 1 stopped
		4 app instances: 2 running, 2 stopped
		0 service instances of type Service Suite
	Space docs-prod is consuming 512 MB memory (0%) of org quota.
		3 apps: 1 running 2 stopped
		5 app instances: 2 running, 3 stopped
		0 service instances of type Service Suite
	Space guillermo-playground is consuming 2560 MB memory (2%) of org quota.
		1 apps: 1 running 0 stopped
		5 app instances: 5 running, 0 stopped
		0 service instances of type Service Suite
	Space haydon-playground is consuming 1024 MB memory (1%) of org quota.
		1 apps: 1 running 0 stopped
		1 app instances: 1 running, 0 stopped
		0 service instances of type Service Suite
	Space jkruck-playground is consuming 128 MB memory (0%) of org quota.
		1 apps: 1 running 0 stopped
		1 app instances: 1 running, 0 stopped
		0 service instances of type Service Suite
	Space rsalas-dev is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space shekel-dev is consuming 1536 MB memory (1%) of org quota.
		3 apps: 3 running 0 stopped
		3 app instances: 3 running, 0 stopped
		0 service instances of type Service Suite
	Space shekel-qa is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space hd-playground is consuming 0 MB memory (0%) of org quota.
		0 apps: 0 running 0 stopped
		0 app instances: 0 running, 0 stopped
		0 service instances of type Service Suite
	Space dwallraff-dev is consuming 1024 MB memory (1%) of org quota.
		1 apps: 1 running 0 stopped
		1 app instances: 1 running, 0 stopped
		0 service instances of type Service Suite
You have deployed 18 apps across 3 org(s), with a total of 27 app instances configured. You are currently running 11 apps with 18 app instances and using 1 service instances of type Service Suite.
```

CSV output:

```
➜  usagereport-plugin git:(master) ✗ cf usage-report -f csv
OrgName, SpaceName, SpaceMemoryUsed, OrgMemoryQuota, AppsDeployed, AppsRunning, AppInstancesDeployed, AppInstancesRunning, ServiceInstanceDeployed
test-org, test-space, 256, 4096, 2, 1, 3, 2, 0
```

##Installation
#####Install from Source (need to have [Go](http://golang.org/dl/) installed)
  ```
  $ go get github.com/cloudfoundry/cli
  $ go get github.com/jigsheth57/trueupreport-plugin
  $ cd $GOPATH/src/github.com/jigsheth57/trueupreport-plugin
  $ go build
  $ cf install-plugin trueupreport-plugin
  ```
