package main

import (
	"fmt"

	mypackage "github.com/matheus0214/myFirstGoProject/myPackage"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(mypackage.Bar)
	mypackage.PrintMine()

	res, rem := division(10, 3)
	fmt.Println(res, rem)

	a := sum(2)(3)
	fmt.Println(a)

	fmt.Println(sumTotal(10, 10, 2))
}

// Using naked return
func division(a, b int) (res int, rem int) {
	res = a / b
	rem = a % b

	return
}

func sum(a int) func(int) int {
	return func(b int) int {
		return a + b
	}
}

func sumTotal(nums ...int) int {
	var out int

	for _, v := range nums {
		out += v
	}

	return out
}
