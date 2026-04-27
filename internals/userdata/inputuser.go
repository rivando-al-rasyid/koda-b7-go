package userdata

import (
	"bufio"
	"fmt"
	"strconv"
	"strings"

	"github.com/rivando-al-rasyid/koda-b7-go/internals/models"
)

func Inputuser(scanner *bufio.Scanner) models.DataDiri {
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

	return models.DataDiri{
		Nama:        nama,
		Foto:        foto,
		Email:       email,
		Umur:        umur,
		NoTel:       noTel,
		SPernikahan: sPernikahan,
		Pendidikan:  pendidikan,
	}
}
