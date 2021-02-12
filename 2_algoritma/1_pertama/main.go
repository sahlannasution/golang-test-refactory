package main

import (
	"fmt"
	"strconv"
)

func main() {
	angka := 1234567
	count := len(strconv.Itoa(angka))
	k := count - 1
	for i := 1; i <= count; i++ {
		result := i
		for j := 0; j < k; j++ {
			result *= 10
		}
		fmt.Println(result)
		k--
	}
}
