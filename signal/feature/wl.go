package feature

import (
	"math"

	"github.com/balazshorvath/go-signal-processor/signal"
)

// NewWLSignalProcessor returns a signal processor, that calculates the Waveform Length or Wave Length
func NewWLSignalProcessor() signal.Processor {
	return func(window []float64) []float64 {
		wl := 0.0
		for i := 0; i < len(window)-1; i++ {
			wl += math.Abs(window[i+1] - window[i])
		}
		return []float64{wl}
	}
}
