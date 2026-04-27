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
func rightAngleTriangle(j int) {
	for i := 0; i <= j; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}

func printCircleResult(r float32) {
	fmt.Print("Keliling Lingkaran : ")
	fmt.Println(circumference(r))
	fmt.Print("Luas Lingkaran : ")
	fmt.Println(circleArea(r))

}

type dataSekolah struct {
	Nama    string
	Jurusan string
}
type dataDiri struct {
	Nama        string
	Foto        string
	Email       string
	Umur        uint8
	NoTel       string
	SPernikahan bool
	Pendidikan  []dataSekolah
}

func main() {
	// var radius float32
	// var base int

	// fmt.Print("Enter a radius circle: ")
	// fmt.Scan(&radius) // Reads input and stores it in radius
	// printCircleResult(radius)
	// fmt.Print("Enter a base of right-angled triangle : ")
	// fmt.Scan(&base) // Reads input and stores it in base
	// rightAngleTriangle(base)

	// numbers := []int{50, 75, 66, 20, 32, 90}
	// front := numbers[0 : len(numbers)-3]
	// end := numbers[len(numbers)-3:]
	// result := append(front, 88)
	// result = append(result, end...)

	user := dataDiri{
		Nama:        "Alex Riva",
		Foto:        "profile.jpg",
		Email:       "alex@example.com",
		Umur:        25,
		NoTel:       "08123456789",
		SPernikahan: false,
		Pendidikan: []dataSekolah{
			{Nama: "SMA Negeri 1", Jurusan: "IPA"},
			{Nama: "Universitas Indonesia", Jurusan: "Teknik Informatika"},
		},
	}
	fmt.Println("User Name:", user.Nama)
	fmt.Println("Major at last school:", user.Pendidikan[1].Jurusan)
}
