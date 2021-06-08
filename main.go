package main

import (
	"fmt"
	_ "learngo/simpleinterest"
	"log"
)

var p, r, t = 5000.0, 10.0, 1.0

func init() {
	println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}

func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println(num, "is even")
		return
	}

	fmt.Println(num, "is odd")
}
