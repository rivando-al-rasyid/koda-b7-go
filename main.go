package main

import (
	"fmt"
	"math"
	"strings"
)

func circleArea(r float32) float32 {
	return math.Pi * r * r
}
func circumference(r float32) float32 {
	const diameter = float32(2) * math.Pi
	return diameter * r
}
func segitiga(j int) {
	for i := 0; i <= j; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

func output(r float32) {
	fmt.Print("Keliling Lingkaran : ")
	fmt.Println(circumference(r))
	fmt.Print("Luas Lingkaran : ")
	fmt.Println(circleArea(r))

}

func main() {
	var radius float32
	var num int

	fmt.Print("Enter a radius: ")
	fmt.Scan(&radius) // Reads input and stores it in num
	output(radius)
	fmt.Print("Enter a number : ")
	fmt.Scan(&num) // Reads input and stores it in num
	segitiga(num)
}
