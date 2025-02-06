package mypackage

import (
	"fmt"

	"github.com/matheus0214/myFirstGoProject/myPackage/internal/foo"
)

var Bar string = "Bar"

func PrintMine() {
	fmt.Println(foo.Mine)
}
