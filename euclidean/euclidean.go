package main

import (
	"flag"
	"fmt"
	"strconv"
)

//this func find greatest common divisor from two argument values

func main() {
	flag.Parse()
	value1 := flag.Arg(0)
	value2 := flag.Arg(1)
	var greaterValue int
	var tinyValue int

	if value1 < value2 {
		greaterValue, _ = strconv.Atoi(value2)
		tinyValue, _ = strconv.Atoi(value1)
	} else {
		greaterValue, _ = strconv.Atoi(value1)
		tinyValue, _ = strconv.Atoi(value2)
	}

	FindCommonDivisor(greaterValue, tinyValue)

}

func FindCommonDivisor(greaterValue int, tinyValue int) {

	reminder := greaterValue % tinyValue
	if reminder == 0 {
		fmt.Printf("最大公約数は%dです\n", tinyValue)
		return
	}

	FindCommonDivisor(tinyValue, reminder)

}
