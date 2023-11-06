

# Overview

- Use [`goph`](https://github.com/melbahja/goph/) to SSH into a Cisco BGP router, run `show ip bgp 4.2.2.2 bestpath`, and parse fields from the IPv4 BGP bestpath entry similar to the one shown below:

```
BGP routing table entry for 4.0.0.0/9, version 154410995
Paths: (21 available, best #7, table default)
  Not advertised to any peer
  Refresh Epoch 1
  3356
    4.68.4.46 from 4.68.4.46 (4.69.184.201)
      Origin IGP, metric 0, localpref 100, valid, external, atomic-aggregate, best
      Community: 3356:0 3356:3 3356:100 3356:123 3356:575 3356:2012
      path 7F2C230B7E50 RPKI State not found
      rx pathid: 0, tx pathid: 0x0
```

- In this case, you can ssh into the router as `rviews` with no password, and thus the reason for using `goph.Password("")`:

```go
// route-views username 'rviews' with no password
client, err := goph.New("rviews", "route-views.routeviews.org", goph.Password(""))
if err != nil {
        logoru.Critical(err.Error())
}
defer client.Close()
```

- This is only an example.  If you only want bgp bestpath for an IP, you don't really need to use [goph](https://github.com/melbahja/goph/), you can get the same information simpler by using [iptoasn](https://github.com/jamesog/iptoasn/).

# Build and Usage

- Run `make all`
- Run `./routeviews_go`

`routeviews_go` will ssh into the router, parse output with [gotextfsm](https://github.com/sirikothe/gotextfsm/), and dump the dictionary of parsed values to `stdout`.

```
$ ./routeviews_go
2023-10-07 15:56:37.920 |  INFO    | map[bgpAsPath:3356 bgpIpv4NextHop:4.68.4.46 bgpIpv4Prefix:4.0.0.0 bgpPrefixLength:9 bgpTableVersion:154410995]

$
```

# Dependencies

- [goph](https://github.com/melbahja/goph/)
- [gotextfsm](https://github.com/sirikothe/gotextfsm/)
- [logoru](https://github.com/gleich/logoru/)

# FAQ

- Q: Why not use [go-ansible](https://github.com/apenella/go-ansible)?  A: This repo is simpler for my use-case.
- Q: Is [goph](https://github.com/melbahja/goph) faster than [python Ansible](https://github.com/ansible/ansible)?  A: I think it's debatable; most of the run-time is waiting on `ssh` / `route-views` itself, and that is a significant portion of either runtime.  That said, I like the portability of the Go runtime-build-system more than `python` / `ansible`; Go builds static-linked binaries for the CPU target.
- Q: Is [goph](https://github.com/melbahja/goph) a replacement for [go-expect](https://github.com/Netflix/go-expect)?  A: Not at this time.

