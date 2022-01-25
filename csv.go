package main

import "io"

func sum(data []float64) float64 {
	sum := 0.0

	for _, v := range data {
		sum += v
	}
	return sum
}

func avg(data []float64) float64 {
	return sum(data) / float64(len(data))
}

// io.Reader: source of the csv file, also helps with testing
// column:
func csv2float(r io.Reader, column int) ([]float64, error) {

}
