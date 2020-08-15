package types

type PingResponse struct {
	Latency  uint           `json:"latency"`
	Online   int            `json:"online"`
	Max      int            `json:"max"`
	Protocol int            `json:"protocol"`
	Favicon  []byte         `json:"favicon"`
	Motd     string         `json:"motd"`
	Server   string         `json:"server"`
	Version  string         `json:"version"`
	Sample   []PlayerSample `json:"sample"`
}
