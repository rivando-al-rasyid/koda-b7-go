package main

import (
	"fmt"
	"math"
)

func circleArea(r float32) float32 {
	return math.Pi * r * r
}
func circumference(r float32) float32 {
	const diameter = float32(2) * math.Pi
	return diameter * r
}
func output(r float32) {
	fmt.Print("Keliling Lingkaran : ")
	fmt.Println(circumference(r))
	fmt.Print("Luas Lingkaran : ")
	fmt.Println(circleArea(r))

}

func main() {
	output(2)
}
