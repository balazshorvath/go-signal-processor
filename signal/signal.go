package signal

// Signal is a representation of a signal with a pre-defined window size.
// A Signal uses Processors to calculate features of itself.
type Signal struct {
	Id         string
	processors []Processor
	window     []float64
	windowSize int
}

// Processor is a function, that can be used to calculate features of a signal.
//
// The result of a Processor might be one, or multiple values depending on the type of work it does.
// For example a DWT, or AR function might want to return multiple coefficients.
// A Processor could implement a 'memory' to propagate values (see RMS implementation for example).
type Processor func(window []float64) []float64

// ProcessorResult is a container for the results of a data push into a signal.
type ProcessorResult struct {
	Id               string      `json:"id"`
	Value            float64     `json:"value"`
	ProcessorResults [][]float64 `json:"processorResults"`
}

// NewSignal creates a new signal with window size and an identifier.
func NewSignal(windowSize int, id string) *Signal {
	return &Signal{
		Id:         id,
		processors: []Processor{},
		window:     make([]float64, windowSize),
		windowSize: windowSize,
	}
}

// AddProcessor adds a signal processor to be executed when pushing a new data point to the window.
// Results will be in the same order as the processors have been appended to the signal.
func (s *Signal) AddProcessor(processor Processor) (index int) {
	index = len(s.processors)
	s.processors = append(s.processors, processor)
	return index
}

// ClearWindow sets the window values to zero.
func (s *Signal) ClearWindow() {
	for i := 0; i < s.windowSize; i++ {
		s.window[i] = 0.0
	}
}

// Push appends the value and shifts the Window.
// Returns the updated feature of the signal for the current Window.
func (s *Signal) Push(value float64) (dataPoint *ProcessorResult) {
	s.window = append(s.window[1:], value)
	results := make([][]float64, len(s.processors))
	for i, v := range s.processors {
		results[i] = v(s.window)
	}
	return &ProcessorResult{
		Id:               s.Id,
		Value:            value,
		ProcessorResults: results,
	}
}
