package main

import (
	"fmt"

	mypackage "github.com/matheus0214/myFirstGoProject/myPackage"
)

func main() {
	fmt.Println("Hello World!")
	fmt.Println(mypackage.Bar)
	mypackage.PrintMine()
}
