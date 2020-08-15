package types

import "time"

type Pinger interface {
	Ping(host string, port uint16) (*PingResponse, error)
	PingWithTimeout(host string, port uint16, timeout time.Duration) (*PingResponse, error)
}
