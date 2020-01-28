# Go signal processor
The goal of this library is to provide a convenient way to process time series.
A signal may have multiple processors, that are applied to a specified window size.

Signal features implemented:
 - Average (Avg)
 - Mean absolute value (MAV)
 - Root mean square (RMS - with and without memory) 
 - Wilson Amplitude (WAMP)
 - Wave length (WL)
## Examples
### Basic usage
```go
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
		if i % 2 == 1 {
			v = -v
		}
		fmt.Printf("%v\n", s.Push(float64(v)))
	}
}
```
### Processor implementations
For examples turn to the signal/features package.

### DWT Processor
Discrete wavelet transform with the Daubechies 4 wavelet with features of the filtered signal.

The example is a solution to a very specific problem,
but shows an additional way processors may be used.
```go
package main

import (
	"errors"
	"fmt"
	"math"

	"github.com/balazshorvath/go-signal-processor/signal"
)

const (
	sqrtThree = 1.732050808
	// DB4 Scaling coefficients
	db4H0 = (1 + sqrtThree) / (4 * math.Sqrt2)
	db4H1 = (3 + sqrtThree) / (4 * math.Sqrt2)
	db4H2 = (3 - sqrtThree) / (4 * math.Sqrt2)
	db4H3 = (1 - sqrtThree) / (4 * math.Sqrt2)
	// DB4 Wavelet coefficients
	db4G0 = db4H3
	db4G1 = -db4H2
	db4G2 = db4H1
	db4G3 = -db4H0
)

func NewDb4Processor(processors []signal.Processor) signal.Processor {
	return func(window []float64) []float64 {
		result := make([]float64, 0)
		cA1, cD1, err := Db4(window)
		if err != nil {
			fmt.Printf("first Db4 transform failed: %v", err)
            return nil
		}
		// cD2 would already be in the sub 25Hz range, there's no point in going lower, or using this result
		cA2, _, err := Db4(cA1)
		if err != nil {
			fmt.Printf("second Db4 transform failed: %v", err)
            return nil
		}
		for _, processor := range processors {
			result = append(result, processor(cD1)...)
			result = append(result, processor(cA2)...)
		}
		return result
	}
}

func Db4(window []float64) (scale []float64, wavelet []float64, error error) {
	windowLength := len(window)
	if windowLength < 4 {
		return nil, nil, errors.New("window is too small")
	}

	halfLength := windowLength / 2
	scale = make([]float64, halfLength)
	wavelet = make([]float64, halfLength)

	i, j := 0, 0
	for ; i < windowLength-3; i, j = i+2, j+1 {
		scale[j] = window[i]*db4H0 + window[i+1]*db4H1 + window[i+2]*db4H2 + window[i+3]*db4H3
		wavelet[j] = window[i]*db4G0 + window[i+1]*db4G1 + window[i+2]*db4G2 + window[i+3]*db4G3
	}
	scale[j] = window[windowLength-2]*db4H0 + window[windowLength-1]*db4H1 + window[0]*db4H2 + window[1]*db4H3
	wavelet[j] = window[windowLength-2]*db4G0 + window[windowLength-1]*db4G1 + window[0]*db4G2 +
		window[1]*db4G3
	return
}
```
