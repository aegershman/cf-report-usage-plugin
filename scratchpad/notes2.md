# notes2

wow what an incredible file name, right?

Before including scs 3.x labels:

```txt
Space scratchpad is consuming 8768 MB memory (42%) of org quota.
        8 apps: 6 running 2 stopped
        13 app instances: 11 running, 2 stopped
        4 service instances of type Service Suite
```

after changing `string.Contains` logic:

```txt
Space scratchpad is consuming 10816 MB memory (52%) of org quota.
        8 apps: 6 running 2 stopped
        15 app instances: 13 running, 2 stopped
        4 service instances of type Service Suite
```

So we know the two scs-service-broker's got added.

```txt
cf curl '/v2/spaces/4ba83909-cde1-448e-873f-d53b27391604/summary'
cf curl '/v2/organizations/a297887a-f58f-4266-9ba2-186563eec13a/summary'
```
