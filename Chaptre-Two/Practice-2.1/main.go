package main

import (
	"fmt"
	"github.com/imquanquan/The-Go-Programming-Language-Practices/Chaptre-Two/Practice-2.1/tempconv"
)

func main() {
	k := tempconv.Kelvin(0)
	fmt.Println(tempconv.KToC(k))
}