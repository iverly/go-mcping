package types

type PingResponse struct {
	Latency     uint
	PlayerCount PlayerCount
	Protocol    int
	Favicon     string
	Motd        string
	Version     string
	Sample      []PlayerSample
}

type PlayerCount struct {
	Online int
	Max    int
}
