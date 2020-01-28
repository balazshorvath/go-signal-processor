package feature

import (
	"math"

	"github.com/balazshorvath/go-signal-processor/signal"
)

// NewWAMPSignalProcessor returns a signal processor, that calculates the Wilson Amplitude
// =SUM(f(x(i)-x(i+1))
// f(x)= 	{1, if x > threshold
//			{0, otherwise
func NewWAMPSignalProcessor(threshold float64) signal.Processor {
	return func(window []float64) []float64 {
		wamp := 0.0
		for i := 0; i < len(window)-1; i++ {
			if math.Abs(window[i]-window[i+1]) > threshold {
				wamp++
			}
		}
		return []float64{wamp}
	}
}
