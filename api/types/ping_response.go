package types

type PingResponse struct {
	Latency  uint           `json:"latency"`
	Online   int            `json:"online"`
	Max      int            `json:"max"`
	Protocol int            `json:"protocol"`
	Favicon  string         `json:"favicon"`
	Motd     string         `json:"motd"`
	Version  string         `json:"version"`
	Sample   []PlayerSample `json:"sample"`
}
