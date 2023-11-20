package main

import (
	"fmt"

	"github.com/anscombe/go-statistics/statistics"
)

func main() {
	// Generate statistics using the Anscombe Quartet data
	data := [][]float64{
		{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0, 10.0}, // Dataset I
		{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0, 10.0}, // Dataset II
		{10.0, 8.0, 13.0, 9.0, 11.0, 14.0, 6.0, 4.0, 12.0, 7.0, 5.0, 10.0}, // Dataset III
		{8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 8.0, 19.0, 8.0, 8.0, 8.0, 8.0},      // Dataset IV
	}

	statistics := statistics.CalculateStatistics(data)
	fmt.Println("Statistics:")
	fmt.Println("Mean:", statistics.Mean)
	fmt.Println("Variance:", statistics.Variance)
	fmt.Println("Correlation:", statistics.Correlation)

	// Add your own advanced and sophisticated code here
	// ...

	// Example: Perform additional calculations

	// Calculate the Standard Deviation for each dataset
	for i, dataset := range data {
		stdDev := statistics.CalculateStandardDeviation(dataset)
		fmt.Printf("Standard Deviation for Dataset %d: %f\n", i+1, stdDev)
	}

	// Calculate the Covariance between two datasets
	covariance := statistics.CalculateCovariance(data[0], data[1])
	fmt.Println("Covariance between Dataset I and Dataset II:", covariance)

	// Calculate the Regression Line for a dataset
	slope, intercept := statistics.CalculateRegressionLine(data[0], data[1])
	fmt.Printf("Regression Line for Dataset I and Dataset II: y = %f*x + %f\n", slope, intercept)

	// Perform a Hypothesis Test for the correlation coefficient
	alpha := 0.05 // significance level
	hypothesisTest := statistics.PerformHypothesisTest(data[0], data[1], alpha)
	fmt.Printf("Hypothesis Test Result for Dataset I and Dataset II: p-value = %f\n", hypothesisTest.PValue)

	// Calculate the Confidence Interval for the mean of a dataset
	confidenceInterval := statistics.CalculateConfidenceInterval(data[0], 0.95) // 95% confidence level
	fmt.Printf("Confidence Interval for Dataset I (95%% confidence level): [%f, %f]\n", confidenceInterval.Lower, confidenceInterval.Upper)
}
