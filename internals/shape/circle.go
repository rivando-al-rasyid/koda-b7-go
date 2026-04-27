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
	fmt.Print("Keliling Lingkaran : ")
	fmt.Println(circumference(r))
	fmt.Print("Luas Lingkaran : ")
	fmt.Println(circleArea(r))
}
