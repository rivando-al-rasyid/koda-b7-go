package shape

import (
	"fmt"
	"strings"
)

func RightAngleTriangle(j int) {
	for i := 0; i <= j; i++ {
		fmt.Println(strings.Repeat("*", i))
	}
}
