package main

import (
	"fmt"
	"github.com/imquanquan/The-Go-Programming-Language-Practices/Chaptre-Two/Practice-2.2/qualityconv"
	"os"
	"strconv"
)

func Usage() {
	fmt.Println("Usage: conv PToK/KToP number")
}

func main()  {
	if len(os.Args[1:]) < 2 || len(os.Args[1:]) > 2 {
		Usage()
		os.Exit(1)
	}
	convType := os.Args[1]
	number, err := strconv.ParseFloat(os.Args[2], 64)
	if err != nil {
		Usage()
		os.Exit(2)
	}
	switch convType {
	case "KToP":
		fmt.Println(qualityconv.KToP(qualityconv.Kilogram(number)))
	case "PToK":
		fmt.Println(qualityconv.PToK(qualityconv.Pound(number)))
	default:
		Usage(); os.Exit(3)
	}
}
