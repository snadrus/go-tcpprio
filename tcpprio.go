// The constants: https://github.com/leostratus/netinet/blob/master/ip.h#L86
// What they mean: https://www.cisco.com/c/en/us/support/docs/quality-of-service-qos/qos-packet-marking/10103-dscpvalues.html
// Originating Doc: https://www.ietf.org/rfc/rfc2597.txt
package tcpprio

import (
	"net"
	"net/http"
	"syscall"
)

type Priority int

const IPTOS_THROUGHPUT Priority = 0x08

// IPTOS_DSCP_AF43 yields the most to other traffic. 5-10% packet loss.
const IPTOS_DSCP_AF43 Priority = 0x98

// IPTOS_DSCP_AF11 is for the most important traffic.
const IPTOS_DSCP_AF11 = 0x28

// IPTOS_DSCP_AF12 is for very important traffic.
const IPTOS_DSCP_AF12 Priority = 0x30

// HttpClient. Defer client.Transport.CloseIdleConnections() when done.
// Use a DSCP priority value.
func HttpClient(prio Priority) *http.Client {
	return &http.Client{
		Transport: Transport(prio),
	}
}

// Transport. Defer transport.CloseIdleConnections() when done.
// Use a DSCP priority value.
func Transport(prio Priority) *http.Transport {
	return &http.Transport{
		DialTLS: func(network, addr string) (net.Conn, error) {
			// Dial the connection with the lowest TCP QoS priority
			netAddr, err := net.ResolveTCPAddr(network, addr)
			if err != nil {
				return nil, err
			}
			conn, err := net.DialTCP(network, nil, netAddr)
			if err != nil {
				return nil, err
			}

			UpdateTCPConn(conn, prio)
			return conn, nil
		},
	}
}

func UpdateTCPConn(conn *net.TCPConn, prio Priority) (err error) {
	// Set the TCP user priority to the lowest value (0x18)
	sys, err := conn.SyscallConn()
	if err != nil {
		return err
	}
	sys.Control(func(fd uintptr) {
		// Set the TCP user priority to the lowest value (0x18)
		err = syscall.SetsockoptInt(int(fd), syscall.IPPROTO_IP, syscall.IP_TOS, int(prio))
	})
	return nil
}
