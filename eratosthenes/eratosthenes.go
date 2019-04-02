package main

import (
	"flag"
	"fmt"
	"strconv"
)

// this func find prime number from　arg value or under

func main() {
	flag.Parse()
	arg := flag.Arg(0)
	argValue, _ := strconv.Atoi(arg)

	arr := make([]int, argValue+1, argValue+1)
	var k int

	for k = 2; k*k <= argValue; k++ {
		var i int
		if arr[k] == 0 {
			for i = k; i <= argValue/k; i++ {
				notPrimeIndex := i * k
				arr[notPrimeIndex] = 1
			}
		}
	}
	var w int
	fmt.Print("Prime number is: ")
	for w = 2; w < argValue; w++ {
		if arr[w] == 0 {
			fmt.Printf(" %d,", w)
		}
	}
	fmt.Println()

}
