package main

import (
	"fmt"
	"os"
	"strconv"
)

func errorArgs() bool {
	i := len(os.Args)

	if i == 2 {
		return true
	}

	return false
}

func convertStringToInt(str string) int {
	nb, _ := strconv.Atoi(str)
	return nb
}

func main() {
	i := os.Args[1]
	nb := convertStringToInt(i)
	errorArgs := errorArgs()

	if errorArgs == false {
		return
	}

	if nb%3 == 0 && nb%5 == 0 {
		fmt.Println("FizzBuzz")
	} else if nb%3 == 0 {
		fmt.Println("Fizz")
	} else if nb%5 == 0 {
		fmt.Println("Buzz")
	} else {
		fmt.Println("GoodBye FizzBuzz")
	}

	return
}
