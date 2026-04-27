package shape

import (
	"fmt"
	"math"
)

func circleArea(r float64) float64 {
	return math.Pi * r * r
}

func circumference(r float64) float64 {
	const diameter = float64(2) * math.Pi
	return diameter * r
}

func PrintCircleResult(r float64) {
	fmt.Printf("Keliling Lingkaran : %.2f\n", circumference(r))
	fmt.Printf("Luas Lingkaran : %.2f\n", circleArea(r))
}
