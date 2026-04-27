package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rivando-al-rasyid/koda-b7-go/internals/codeslices"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/models"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/shape"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("What do you want?\n1. Circumference and Circle Area\n2. Star Right Angle Triangle\n3. Go Slice Manipulation\n4. Register\nType a number: ")

	if scanner.Scan() {
		choice := scanner.Text()

		switch choice {
		case "1":
			fmt.Print("Enter radius: ")
			scanner.Scan()
			radius, err := strconv.ParseFloat(strings.TrimSpace(scanner.Text()), 64)
			if err != nil {
				fmt.Println("Invalid radius:", err)
				return
			}
			shape.PrintCircleResult(radius)

		case "2":
			fmt.Print("Enter base: ")
			scanner.Scan()
			base, err := strconv.Atoi(strings.TrimSpace(scanner.Text()))
			if err != nil {
				fmt.Println("Invalid base:", err)
				return
			}
			shape.RightAngleTriangle(base)

		case "3":
			codeslices.SliceManipulation()

		case "4":
			fmt.Print("Masukkan Nama: ")
			scanner.Scan()
			nama := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan Foto (filename): ")
			scanner.Scan()
			foto := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan Email: ")
			scanner.Scan()
			email := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan Umur: ")
			scanner.Scan()
			umur, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			fmt.Print("Masukkan No Telepon: ")
			scanner.Scan()
			noTel := strings.TrimSpace(scanner.Text())

			fmt.Print("Sudah Menikah? (true/false): ")
			scanner.Scan()
			sPernikahan, _ := strconv.ParseBool(strings.TrimSpace(scanner.Text()))

			fmt.Print("Masukkan jumlah pendidikan: ")
			scanner.Scan()
			jumlahPendidikan, _ := strconv.Atoi(strings.TrimSpace(scanner.Text()))

			var pendidikan []models.DataSekolah
			for i := 0; i < jumlahPendidikan; i++ {
				fmt.Printf("-- Pendidikan %d --\n", i+1)

				fmt.Print("Nama Sekolah: ")
				scanner.Scan()
				namaSekolah := strings.TrimSpace(scanner.Text())

				fmt.Print("Jurusan: ")
				scanner.Scan()
				jurusan := strings.TrimSpace(scanner.Text())

				pendidikan = append(pendidikan, models.DataSekolah{
					Nama:    namaSekolah,
					Jurusan: jurusan,
				})
			}

			user := models.DataDiri{
				Nama:        nama,
				Foto:        foto,
				Email:       email,
				Umur:        umur,
				NoTel:       noTel,
				SPernikahan: sPernikahan,
				Pendidikan:  pendidikan,
			}

			fmt.Printf("\n=== Data Terdaftar ===\n")
			fmt.Printf("Nama        : %s\n", user.Nama)
			fmt.Printf("Foto        : %s\n", user.Foto)
			fmt.Printf("Email       : %s\n", user.Email)
			fmt.Printf("Umur        : %d\n", user.Umur)
			fmt.Printf("No Telepon  : %s\n", user.NoTel)
			fmt.Printf("Menikah     : %v\n", user.SPernikahan)
			fmt.Printf("Pendidikan  :\n")
			for i, p := range user.Pendidikan {
				fmt.Printf("  %d. %s - %s\n", i+1, p.Nama, p.Jurusan)
			}

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
