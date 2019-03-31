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
	arr := strings.Split(flag.Arg(1), ",")

	switch algo {
	case "selection":
		SelectionSort(arr)
	case "bubble":
		BubbleSort(arr)
	case "insert":
		InsertSort(arr)
	case "quick":
		QuickSort(arr, 0, len(arr)-1)
	}
}

func SelectionSort(arr []string) {
	var i int
	for i = 0; i < len(arr)-1; i++ {
		indexMin := i
		var k int
		for k = i + 1; k < len(arr); k++ {
			arrValue, _ := strconv.Atoi(arr[k])
			indexMinValue, _ := strconv.Atoi(arr[indexMin])
			if arrValue < indexMinValue {
				indexMin = k
			}
		}

		w := arr[i]
		arr[i] = arr[indexMin]
		arr[indexMin] = w

		fmt.Printf("%d番目までソート済み：%#v\n", i+1, arr)
	}

	fmt.Println(arr)

}

func BubbleSort(arr []string) {
	var i int
	var k int
	var w string

	for k = 0; k < len(arr); k++ {
		for i = len(arr) - 1; i > k; i-- {
			leftValue, _ := strconv.Atoi(arr[i-1])
			rightValue, _ := strconv.Atoi(arr[i])

			if leftValue > rightValue {
				w = arr[i-1]
				arr[i-1] = arr[i]
				arr[i] = w
			}

		}

		fmt.Printf("%d番目までソート済み：%#v\n", k+1, arr)

	}

	fmt.Println(arr)
}

func InsertSort(arr []string) {

	var i int
	var k int
	var x string

	for i = 1; i < len(arr); i++ {
		x = arr[i]
		insertValue, _ := strconv.Atoi(x)

		for k = i; k > 0; k-- {
			arrValue, _ := strconv.Atoi(arr[k-1])
			if arrValue > insertValue {
				arr[k] = arr[k-1]
			} else {
				break
			}
		}
		arr[k] = x

		fmt.Printf("%d番目までソート済み：%#v\n", i+1, arr)

	}

	fmt.Println(arr)

}

func QuickSort(arr []string, left int, right int) {
	i := left + 1
	k := right
	var w string
	standardValue, _ := strconv.Atoi(arr[left])

	for {
		if i >= k {
			break
		}
		//get index that value is higher than standard value
		for {
			arrValueFromLeft, _ := strconv.Atoi(arr[i])

			if arrValueFromLeft < standardValue && i < right {
				i = i + 1
				continue
			}
			break
		}

		//get index that value is lower than standart value

		for {
			arrValueFromRight, _ := strconv.Atoi(arr[k])

			if arrValueFromRight >= standardValue && k > left {
				k = k - 1
				continue
			}
			break

		}
		//change array value that higher than standard value from
		//left index and lower than standard value from right index
		if i < k {
			w = arr[i]
			arr[i] = arr[k]
			arr[k] = w
		}

	}

	canterArrValue, _ := strconv.Atoi(arr[k])

	if right-left != 1 || standardValue > canterArrValue {
		w = arr[left]
		arr[left] = arr[k]
		arr[k] = w
	}

	fmt.Println(arr)

	if left < k-1 {
		fmt.Printf("left side index is %d, right side index is %d\n", left, k-1)
		fmt.Println("sort left side")
		QuickSort(arr, left, k-1)
	}

	if k+1 < right {
		fmt.Printf("left side index is %d, right side index is %d\n", k+1, right)
		fmt.Println("sort right side")
		QuickSort(arr, k+1, right)
	}
}
