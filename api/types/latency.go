package types

// An Latency represent an object that can calculate the latency between Start() and End()
type Latency interface {
	Start() // Start the latency
	End() uint64 // Stop the latency and calculate it
	getMS() uint64 // Get the current MS of the system
}
