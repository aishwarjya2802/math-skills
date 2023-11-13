package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
)

func main() {
	var data []float64
	text, _ := os.ReadFile("data.txt")           //reads file
	textstring := (strings.Fields(string(text))) //removes whitespace
	for _, char := range textstring {
		char, _ := strconv.ParseFloat(string(char), 64)
		data = append(data, char)
	}

	//Calculate and print the statistics
	average := math.Round(calculateAverage(data))
	median := math.Round(calculateMedian(data))
	variance := math.Round(calculateVariance(data, average))
	standardDeviation := math.Round(calculateStandardDeviation(variance))

	fmt.Println("Average:", int(average))
	fmt.Println("Median:", median)
	fmt.Println("Variance:", int(variance))
	fmt.Println("Standard Deviation:", int(standardDeviation))

}

// Calculates the average of a slice of float64 numbers
func calculateAverage(data []float64) float64 {
	sum := 0.0
	for _, num := range data {
		sum += num
	}
	return sum / float64(len(data))
}

// Calculates the median of a slice of float64 numbers.
func calculateMedian(data []float64) float64 {
	sort.Float64s(data)
	n := len(data)
	if n%2 == 0 {
		// If the number of elements is even, take the average of the middle two values.
		mid := n / 2
		return (data[mid-1] + data[mid]) / 2.0
	}
	// If the number of elements is odd
	return data[n/2]
}

// Calculates the variance of a slice of float64 numbers.
func calculateVariance(data []float64, mean float64) float64 {
	sum := 0.0
	for _, num := range data {
		deviation := num - mean
		sum += deviation * deviation
	}
	return sum / float64(len(data))
}

// Calculates the standard deviation from the variance.
func calculateStandardDeviation(variance float64) float64 {
	return math.Sqrt(variance)
}
