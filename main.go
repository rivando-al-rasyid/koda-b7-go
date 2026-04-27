package main

import "fmt"

func main() {
	// Circle
	var radius float32
	fmt.Print("Enter a radius circle: ")
	fmt.Scan(&radius)
	printCircleResult(radius)

	// Triangle
	var base int
	fmt.Print("Enter a base of right-angled triangle : ")
	fmt.Scan(&base)
	rightAngleTriangle(base)

	slices()
	// User profile
	user := dataDiri{
		Nama:        "vando",
		Foto:        "profile.jpg",
		Email:       "vando@example.com",
		Umur:        25,
		NoTel:       "08123456789",
		SPernikahan: false,
		Pendidikan: []dataSekolah{
			{Nama: "SMA Negeri 1", Jurusan: "IPA"},
			{Nama: "Binus", Jurusan: "Magister (S2) Sistem Informasi"},
		},
	}
	fmt.Println("User Name:", user.Nama)
	fmt.Println("Major at last school:", user.Pendidikan[1].Jurusan)
}
