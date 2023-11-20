package main

import (
	"fmt"

	"github.com/example/original_code/statistics"
)

func main() {
	data := []float64{10, 8, 13, 9, 11, 14, 6, 4, 12, 7, 5}

	// Calculate mean
	mean := statistics.CalculateMean(data)

	// Calculate variance
	variance := statistics.CalculateVariance(data)

	// Calculate standard deviation
	standardDeviation := statistics.CalculateStandardDeviation(data)

	// Calculate median
	median := statistics.CalculateMedian(data)

	// Calculate mode
	mode := statistics.CalculateMode(data)

	// Calculate range
	min, max := statistics.CalculateRange(data)

	fmt.Println("Mean:", mean)
	fmt.Println("Variance:", variance)
	fmt.Println("Standard Deviation:", standardDeviation)
	fmt.Println("Median:", median)
	fmt.Println("Mode:", mode)
	fmt.Println("Range:", min, "-", max)
}
