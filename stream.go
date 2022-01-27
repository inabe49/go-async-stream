package stream

import (
	"time"
)


type Stream interface {
	Next(timeout time.Duration) (interface{}, bool)
}

type StreamSource struct {
}

// GetStream makes new Stream.
func (s *StreamSource) GetStream(capacity int) Stream {
	return &stream{}
}

// Publish broadcast data for every Streams
func (s *StreamSource) Publish(data interface{}) {

}

type stream struct {
}

func (s *stream) Next(timeout time.Duration) (interface{}, bool) {
	return 1, true
}
