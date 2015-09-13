package main

import (
	"fmt"
	"strconv"
)

func main() {
	for i := 1; i <= 100; i++ {
		fmt.Println(foobarqix(i))
	}
}

func foobarqix(n int) string {
	result := ""

	if n%3 == 0 {
		result += "Foo"
	}
	if n%5 == 0 {
		result += "Bar"
	}
	if n%7 == 0 {
		result += "Qix"
	}

	digits := strconv.Itoa(n)

	for _, digit := range digits {
		switch digit {
		case '3':
			result += "Foo"
		case '5':
			result += "Bar"
		case '7':
			result += "Qix"
		}
	}

	if result == "" {
		result += digits
	}

	return result
}
