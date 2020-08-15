package types

// An PingResponse represent the structure returned by the server
type PingResponse struct {
	Latency     uint // Latency between you and the server
	PlayerCount PlayerCount // Players count information of the server
	Protocol    int // Protocol number of the server
	Favicon     string // Favicon in base64 of the server
	Motd        string // Motd of the server without color
	Version     string // Version of the server
	Sample      []PlayerSample // List of connected players on the server
}

// An PlayerCount represent the player count information of an server
type PlayerCount struct {
	Online int // Number of connected players on the server
	Max    int // Number of maximum players on the server
}
