package feature

import (
	"math"

	"github.com/balazshorvath/go-signal-processor/signal"
)

// rms is the memory for the RMS processor
type rms struct {
	previousRms float64
	count       uint64
}

// NewRMSProcessor returns a signal processor, that calculates the root mean square of the whole signal
func NewRMSProcessor() signal.Processor {
	memory := rms{
		previousRms: 0.0,
		count:       0,
	}
	return func(window []float64) []float64 {
		memory.count++
		// To avoid overflow, calculate iteratively
		// https://stats.stackexchange.com/questions/221826/is-it-possible-to-compute-rmse-iteratively
		// sqrt((t-1/t)*prev^2 + (curr^2/t))
		newValue := window[len(window)-1]
		memory.previousRms = math.Sqrt(
			((float64(memory.count-1) / float64(memory.count)) * (memory.previousRms * memory.previousRms)) +
				(newValue * newValue / float64(memory.count)),
		)
		return []float64{memory.previousRms}
	}
}

// NewRMSWindowProcessor returns a signal processor, that calculates the root mean square of the signal window
func NewRMSWindowProcessor() signal.Processor {
	return func(window []float64) []float64 {
		sum := 0.0
		for _, v := range window {
			sum += v * v
		}
		return []float64{math.Sqrt(sum / float64(len(window)))}
	}
}
