package main

import (
	"fmt"
	"strings"
)

func rightAngleTriangle(j int) {
	for i := 0; i <= j; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
