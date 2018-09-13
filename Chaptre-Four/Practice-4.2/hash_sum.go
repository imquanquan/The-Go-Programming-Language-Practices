package main

import (
	"crypto/sha256"
	"crypto/sha512"
	"fmt"
	"os"
)

func main()  {
	var s, bit string
	switch len(os.Args) {
	case 2:
		s, bit = os.Args[1], "256"
	case 3:
		s, bit = os.Args[1], os.Args[2]
	default:
		fmt.Println("Usage: hash_sum string [256/384/512]")
	}
	switch bit {
	case "256":
		fmt.Printf("%x\n", sha256.Sum256([]byte(s)))
	case "384":
		fmt.Printf("%x\n", sha512.Sum384([]byte(s)))
	case "512":
		fmt.Printf("%x\n", sha512.Sum512([]byte(s)))
	default:
		fmt.Println("Usage: hash_sum string [256/384/512]")
	}
}
