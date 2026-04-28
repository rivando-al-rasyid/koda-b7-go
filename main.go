package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rivando-al-rasyid/koda-b7-go/internals/bank"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/codeslices"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/localfile"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/models"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/shape"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/userdata"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Printf("What do you want?\n1. Circumference and Circle Area\n2. Star Right Angle Triangle\n3. Go Slice Manipulation\n4. Register\nType a number: \n5. read File\n6. greeting \nType a number:")

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
			userData := userdata.Inputuser(scanner)
			userdata.Showuser(userData)
		case "5":
			localfile.Inputfile()

		case "6":
			fmt.Print("Masukkan Nama: ")
			scanner.Scan()

			name := strings.TrimSpace(scanner.Text())
			fmt.Print("Alamat : ")
			scanner.Scan()

			address := strings.TrimSpace(scanner.Text())

			fmt.Print("Masukkan Email: ")
			scanner.Scan()
			email := strings.TrimSpace(scanner.Text())

			p1 := models.AddPerson(name, address, email)
			p1.Greet()
			p1.SetPerson("mayum")
			p1.Greet()

		case "7":
			myTransfer := bank.Transfer{Jumlah: 100}
			fmt.Println("Transfer amount:", myTransfer.Transaction())

		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}

}
