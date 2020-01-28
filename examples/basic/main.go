package main

import (
	"fmt"

	"github.com/balazshorvath/go-signal-processor/signal"
	"github.com/balazshorvath/go-signal-processor/signal/feature"
)

func main() {
	// Create a signal
	s := signal.NewSignal(10, "zero-to-hundred")
	// Add processors
	s.AddProcessor(feature.NewMAVSignalProcessor())
	s.AddProcessor(feature.NewWLSignalProcessor())
	// Push values and retrieve the results
	for i := 0; i <= 100; i++ {
		v := i
		if i%2 == 1 {
			v = -v
		}
		fmt.Printf("%v\n", s.Push(float64(v)))
	}
}
