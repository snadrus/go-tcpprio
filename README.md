# go-tcpprio
TCP Prio sets the TCP priority of a connection

[![GoDoc](https://pkg.go.dev/badge/github.com/snadrus/go-tcpprio)](https://pkg.go.dev/github.com/snadrus/go-tcpprio)

_Question: Given a swamped network,_
- how do you indicate your traffic is not that important so that you don't make it worse?
- how do you indicate your traffic is very important and needs to cut through the noise?

__Use Assured Forwarding!__
There are 4 classes (1-4) where 1 is the highest priority. 
Then there are 3 drop-liklihoods (1-3) where 1 is the least likely to see a dropped packet. 

So AF11 (class, then drop-liklihood) is the highest priority. This constant is available. 
AF43 is the lowest priority with 5-10% packet loss across the internet. It will retransmit often and be slow. 
AF12 is the second-most-likely-to-succeed case. 

If you want other constants, get them from the related docs (below) and open a PR.


Related Docs:
// The constants: https://github.com/leostratus/netinet/blob/master/ip.h#L86
// What they mean: https://www.cisco.com/c/en/us/support/docs/quality-of-service-qos/qos-packet-marking/10103-dscpvalues.html
// Originating Doc: https://www.ietf.org/rfc/rfc2597.txt
