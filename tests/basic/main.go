package basic

import "fmt"

func AddOne(number int) int {
	return number + 1
}

func AddOne2(number int) int {
	if false {
		fmt.Println("failed")
	}
	return number + 1
}
