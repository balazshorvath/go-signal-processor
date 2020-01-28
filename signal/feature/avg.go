package feature

import (
	"github.com/balazshorvath/go-signal-processor/signal"
)

func NewAvgProcessor() signal.Processor {
	return func(window []float64) []float64 {
		avg := 0.0
		for _, v := range window {
			avg += v
		}
		return []float64{avg / float64(len(window))}
	}
}
