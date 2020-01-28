package main

import (
	"encoding/csv"
	"fmt"
	"io"
	"os"
	"strconv"

	"github.com/balazshorvath/go-signal-processor/signal"
	"github.com/balazshorvath/go-signal-processor/signal/feature"
)

var (
	signals = make(map[string]*signal.Signal)
)

func main() {
	// https://pkgstore.datahub.io/core/global-temp/annual_csv/data/a26b154688b061cdd04f1df36e4408be/annual_csv.csv
	file, err := os.Open("global-annual-temp.csv")
	if err != nil {
		panic(err)
	}
	fmt.Println("Source\tYear\tValue\tRMS\tWL")
	reader := csv.NewReader(file)
	for {
		record, err := reader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			panic(err)
		}
		s := getSignal(record[0])
		avg, _ := strconv.ParseFloat(record[2], 64)
		result := s.Push(avg)
		fmt.Printf(
			"%s\t%s\t%.2f\t%.2f\t%.2f\n",
			record[0],
			record[1],
			result.Value,
			result.ProcessorResults[0][0],
			result.ProcessorResults[1][0],
		)
	}
}

func getSignal(id string) *signal.Signal {
	s, ok := signals[id]
	if !ok {
		s = signal.NewSignal(10, id)
		s.AddProcessor(feature.NewRMSProcessor())
		s.AddProcessor(feature.NewWLSignalProcessor())
		signals[id] = s
	}
	return s
}
