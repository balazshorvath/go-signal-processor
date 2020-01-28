package feature

import (
	"math"

	"github.com/balazshorvath/go-signal-processor/signal"
)

// NewMAVSignalProcessor returns a signal processor, that calculates the Mean Absolute Value
func NewMAVSignalProcessor() signal.Processor {
	return func(window []float64) []float64 {
		mav := 0.0
		for _, v := range window {
			mav += math.Abs(v)
		}
		return []float64{mav / float64(len(window))}
	}
}
