package main

import (
	"fmt"

	"github.com/rivando-al-rasyid/koda-b7-go/internals/codeslices"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/models"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/shape"
)

func main() {
	// Circle
	var radius float32
	fmt.Print("Enter a radius circle: ")
	fmt.Scan(&radius)
	shape.PrintCircleResult(radius)

	// Triangle
	var base int
	fmt.Print("Enter a base of right-angled triangle : ")
	fmt.Scan(&base)
	shape.RightAngleTriangle(base)

	codeslices.SliceManipulation()
	// User profile
	user := models.DataDiri{
		Nama:        "vando",
		Foto:        "profile.jpg",
		Email:       "vando@example.com",
		Umur:        25,
		NoTel:       "08123456789",
		SPernikahan: false,
		Pendidikan: []models.DataSekolah{ // ← models.DataSekolah
			{Nama: "SMA Negeri 1", Jurusan: "IPA"},
			{Nama: "Binus", Jurusan: "Magister (S2) Sistem Informasi"},
		},
	}
	fmt.Println("User Name:", user.Nama)
	fmt.Println("Major at last school:", user.Pendidikan[1].Jurusan)
}
