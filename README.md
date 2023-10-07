

# Overview

- Use [goph](https://github.com/melbahja/goph/) to SSH into a Cisco BGP router, run `show ip bgp 4.2.2.2 bestpath`, and parse fields from the IPv4 BGP bestpath entry similar to the one shown below:

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

- This is only for an example.  If you only want bgp bestpath for an IP, you don't really need to use `goph`, you can get the same information much faster by using [iptoasn](https://github.com/jamesog/iptoasn/).
