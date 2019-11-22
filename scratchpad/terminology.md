# terminology

will need to make sure I'm using official terminology in the future, but let me get this straight at least for myself:

## lApps

"length of apps in a space" == unique app guids, including stopped ones. Does NOT capture the app _instances_ of the app guids

so if you had this from `cf apps`:

```txt
hammerdb-test                   stopped           0/1
nodejs-web                      started           2/2
push-test-webhook-switchboard   started           2/2
```

that would be 3 apps, even though one of them is `stopped`

## rApps

"running apps" == unique app guids, but however many of them are actually running

so this:

```txt
hammerdb-test                   stopped           0/1
nodejs-web                      started           2/2
push-test-webhook-switchboard   started           2/2
```

would be 2

## InstancesCount

"total number of app AIs" == unique app guids multiplied by however many instances they have declared

started/stopped doesn't matter

```txt
hammerdb-test                   stopped           0/1
nodejs-web                      started           2/2
push-test-webhook-switchboard   started           2/2
```

this would be...

```txt
hammerdb-test => 1
nodejs-web => 2
push-test-webhook-switchboard => 2
```

total of 5

## lAIs

here's where things get weird. lAIs is "length of total number of AIs". But first you have to ask, _what is an AI?_ `p-spring-cloud-config` is an SI from CF's perspective, but gets _billed_ as an AI. Currently, "billable AIs" like SCS or SCDF are lumped into lAIs and reported as a "running app instance"

I propose we change this.

We need to distinguish between "billable AIs" and "real AIs", or something to that effect. Maybe call it "cf AIs"? That's lame, dont' call it that.

... `cannonicalAI`?

(bAIs for billableAIs?)

How about an AI is an app within an org/space, and then a "billableAI" is a superset of AIs and _anything else_, like SIs, etc.?

Let me just try `cAIs` == cannonical AI, which is what cf would report as an application running within the org/space

## rAIs

"running AIs", which are a combination of "running cAIs and bAIs"... e.g., SCS, SCDF, etc.
