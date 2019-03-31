package main

import (
	"flag"
	"fmt"
	"strconv"
	"strings"
)

func main() {
	flag.Parse()
	algo := flag.Arg(0)
	searchedValue := flag.Arg(1)
	arr := strings.Split(flag.Arg(2), ",")

	switch algo {
	case "linear":
		LinearSearch(searchedValue, arr)
	case "binary":
		BinarySearch(searchedValue, arr)
	case "hash":
		HashSearch(searchedValue, arr)
	}
}

func LinearSearch(searchedValue string, arr []string) {
	for i := 0; i < len(arr); i++ {
		if searchedValue == arr[i] {
			fmt.Printf("%d番目の値と一致しました", i+1)
			return
		}
	}

	fmt.Println("見つかりませんでした")
	return
}

func BinarySearch(searchedValue string, arr []string) {
	head := 0
	tail := len(arr) - 1

	for {
		if head > tail {
			fmt.Println("一致する値は見つかれませんでした")
			return
		}
		center := (head + tail) / 2
		fmt.Printf("%d番目の値を検索します\n", center+1)
		if arr[center] == searchedValue {
			fmt.Printf("%d番目の値と一致しました\n", center+1)
			return
		}
		if arr[center] < searchedValue {
			head = center + 1
		} else {
			tail = center - 1
		}
	}
}

func HashSearch(searchedValue string, arr []string) {

	hashTable := map[int]int{}

	// this divisor is used at hash func
	divisor := len(arr) * 2

	var i int

	for i = 0; i < len(arr); i++ {
		arrValue, err := strconv.Atoi(arr[i])
		if err != nil {
			fmt.Println(err)
			return
		}
		hashedKey := hash(arrValue, divisor)
		for {
			_, ok := hashTable[hashedKey]
			if !ok {
				hashTable[hashedKey] = arrValue
				break
			} else {
				hashedKey = hash(hashedKey+1, divisor)
			}
		}
	}

	searchValue, err := strconv.Atoi(searchedValue)

	if err != nil {
		fmt.Println(err)
		return
	}

	hashedSearchKey := hash(searchValue, divisor)

	for {
		_, ok := hashTable[hashedSearchKey]
		if !ok {
			fmt.Println("一致する値はありませんでした")
			return
		}

		fmt.Printf("%d番目の値は%dです\n", hashedSearchKey, hashTable[hashedSearchKey])

		if searchValue == hashTable[hashedSearchKey] {
			fmt.Printf("%d番目の値と一致しました\n", hashedSearchKey)
			return
		}

		hashedSearchKey = hash(hashedSearchKey+1, divisor)
	}

}

//function to return hashed value. Hashed Algotithm is going to be changed.
func hash(rawValue int, divisor int) int {
	hashedValue := rawValue % divisor
	return hashedValue
}
