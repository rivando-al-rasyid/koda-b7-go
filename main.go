package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/rivando-al-rasyid/koda-b7-go/internals/codeslices"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/shape"
	"github.com/rivando-al-rasyid/koda-b7-go/internals/userdata"
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
			userData := userdata.Inputuser(scanner)
			userdata.Showuser(userData)
		default:
			fmt.Println("Invalid choice. Please enter 1, 2, 3, or 4.")
		}
	}

	if err := scanner.Err(); err != nil {
		fmt.Println("Error reading input:", err)
	}
}
