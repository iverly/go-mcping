package mcping

import (
	"github.com/iverly/go-mcping/api/types"
	"github.com/iverly/go-mcping/dns"
	"net"
	"strconv"
	"time"
)

type pinger struct {
	DnsResolver types.DnsResolver
}

func NewPinger() *pinger {
	var resolver types.DnsResolver
	resolver = dns.NewResolver()
	return &pinger{DnsResolver: resolver}
}

func NewPingerWithDnsResolver(dnsResolver types.DnsResolver) *pinger {
	return &pinger{DnsResolver: dnsResolver}
}

func (p *pinger) Ping(host string, port uint16) (*types.PingResponse, error) {
	return p.PingWithTimeout(host, port, 3 * time.Second)
}

func (p *pinger) PingWithTimeout(host string, port uint16, timeout time.Duration) (*types.PingResponse, error) {
	resolve, hostSRV, portSRV := p.DnsResolver.SRVResolve(host)
	if resolve {
		host = hostSRV
		port = portSRV
	}

	addr := host + ":" + strconv.Itoa(int(port))
	// Open connection to server
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sendPacket(host, port, &conn)
	response, err := readResponse(&conn)
	if err != nil {
		return nil, err
	}
	return decodeResponse(response), nil
}
