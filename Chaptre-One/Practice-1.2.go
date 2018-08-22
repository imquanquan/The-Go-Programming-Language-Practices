package main

import (
	"os"
	"fmt"
)

func main()  {
	for index, val := range os.Args[1:] {
		fmt.Printf("%d：%s\n", index + 1, val)
	}
}
