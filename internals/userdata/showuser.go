package userdata

import (
	"fmt"

	"github.com/rivando-al-rasyid/koda-b7-go/internals/models"
)

func Showuser(user models.DataDiri) {
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

}
