package statistics

import (
	"math"
	"sort"
)

// CalculateMean calculates the mean of a given slice of float64 values.
func CalculateMean(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum / float64(len(data))
}

// CalculateVariance calculates the variance of a given slice of float64 values.
func CalculateVariance(data []float64) float64 {
	mean := CalculateMean(data)
	sumOfSquares := 0.0
	for _, value := range data {
		sumOfSquares += math.Pow(value-mean, 2)
	}
	return sumOfSquares / float64(len(data))
}

// CalculateStandardDeviation calculates the standard deviation of a given slice of float64 values.
func CalculateStandardDeviation(data []float64) float64 {
	variance := CalculateVariance(data)
	return math.Sqrt(variance)
}

// CalculateMedian calculates the median of a given slice of float64 values.
func CalculateMedian(data []float64) float64 {
	sort.Float64s(data)
	length := len(data)
	if length%2 == 0 {
		return (data[length/2-1] + data[length/2]) / 2
	} else {
		return data[length/2]
	}
}

// CalculateMode calculates the mode(s) of a given slice of float64 values.
func CalculateMode(data []float64) []float64 {
	frequencyMap := make(map[float64]int)
	maxFrequency := 0
	for _, value := range data {
		frequencyMap[value]++
		if frequencyMap[value] > maxFrequency {
			maxFrequency = frequencyMap[value]
		}
	}

	var mode []float64
	for value, frequency := range frequencyMap {
		if frequency == maxFrequency {
			mode = append(mode, value)
		}
	}

	return mode
}

// CalculateRange calculates the range of a given slice of float64 values.
func CalculateRange(data []float64) float64 {
	sort.Float64s(data)
	return data[len(data)-1] - data[0]
}

// CalculateSum calculates the sum of a given slice of float64 values.
func CalculateSum(data []float64) float64 {
	sum := 0.0
	for _, value := range data {
		sum += value
	}
	return sum
}

// CalculateProduct calculates the product of a given slice of float64 values.
func CalculateProduct(data []float64) float64 {
	product := 1.0
	for _, value := range data {
		product *= value
	}
	return product
}

// CalculateMin calculates the minimum value in a given slice of float64 values.
func CalculateMin(data []float64) float64 {
	min := math.Inf(1)
	for _, value := range data {
		if value < min {
			min = value
		}
	}
	return min
}

// CalculateMax calculates the maximum value in a given slice of float64 values.
func CalculateMax(data []float64) float64 {
	max := math.Inf(-1)
	for _, value := range data {
		if value > max {
			max = value
		}
	}
	return max
}
