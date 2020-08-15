package types

import "time"

// An Pinger represent an object that can Ping a Minecraft server
type Pinger interface {
	Ping(host string, port uint16) (*PingResponse, error) // Ping a server with a timeout of 3s
	PingWithTimeout(host string, port uint16, timeout time.Duration) (*PingResponse, error) // Ping a server with a custom timeout
}
