package latency

import (
	"time"
)

type latency struct {
	start   uint64
	end     uint64
	latency uint64
}

func NewLatency() *latency {
	return &latency{}
}

func (l *latency) Latency() uint64 {
	return l.latency
}

func (l *latency) Start() {
	l.start = l.getMS()
}

func (l *latency) End() uint64 {
	l.end = l.getMS()
	l.latency = l.end - l.start
	return l.latency
}

func (l *latency) getMS() uint64 {
	return uint64(time.Now().UnixNano() / int64(time.Millisecond))
}
