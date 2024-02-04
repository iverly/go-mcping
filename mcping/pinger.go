package mcping

import (
	"errors"
	"net"
	"strconv"
	"time"

	"github.com/iverly/go-mcping/api/types"
	"github.com/iverly/go-mcping/dns"
	"github.com/iverly/go-mcping/latency"
)

type Pinger struct {
	DnsResolver types.DnsResolver
	Latency     types.Latency
}

// Create a new Minecraft Pinger
func NewPinger() *Pinger {
	var resolver types.DnsResolver
	resolver = dns.NewResolver()
	return &Pinger{DnsResolver: resolver}
}

// Create a new Minecraft Pinger with a custom DNS resolver
func NewPingerWithDnsResolver(dnsResolver types.DnsResolver) *Pinger {
	return &Pinger{DnsResolver: dnsResolver}
}

// Ping and get information from an host and port, default timeout: 3s
// Return a pointer to types.PingResponse or an error
//
// # Error is thrown when the host is unreachable or the data received are incorrect
//
// Example: Pinger.Ping("play.hypixel.net", 25565)
func (p *Pinger) Ping(host string, port uint16) (*types.PingResponse, error) {
	return p.PingWithTimeout(host, port, 3*time.Second)
}

// Ping and get information from an host and port with a custom timeout
// Return a pointer to types.PingResponse or an error
//
// # Error is thrown when the host is unreachable or the data received are incorrect
//
// Example: pinger.Ping("play.hypixel.net", 25565, 5 * time.Second)
func (p *Pinger) PingWithTimeout(host string, port uint16, timeout time.Duration) (*types.PingResponse, error) {
	resolve, hostSRV, portSRV := p.DnsResolver.SRVResolve(host)
	if resolve {
		host = hostSRV
		port = portSRV
	}

	lat := latency.NewLatency()
	lat.Start()
	addr := host + ":" + strconv.Itoa(int(port))

	// Open connection to server
	conn, err := net.DialTimeout("tcp", addr, timeout)
	if err != nil {
		return nil, err
	}
	defer conn.Close()

	sendPacket(host, port, &conn)
	type Response struct {
		Response string
		Error    error
	}
	result := make(chan Response, 1)
	go func() {
		response, err := readResponse(&conn)
		result <- Response{
			Response: response,
			Error:    err,
		}
	}()
	select {
	case <-time.After(timeout):
		return nil, errors.New("timed out while reading server response")
	case result := <-result:
		if result.Error != nil {
			return nil, result.Error
		}

		lat.End()
		decode := decodeResponse(result.Response)
		decode.Latency = uint(lat.Latency())
		return decode, nil
	}
}
