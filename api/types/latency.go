package types

type Latency interface {
	Start()
	End() uint64
	getMS() uint64
}
