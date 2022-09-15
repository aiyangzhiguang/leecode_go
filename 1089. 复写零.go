package main

import "fmt"

func duplicateZeros(arr []int) {
	//n := len(arr)
	for i, x := range arr {
		if x == 0 {
			arr = append(append(arr[:i], 0), arr[i:]...)
		}
	}
	fmt.Println(arr)
	println(&arr)
	//arr = arr[:n]
	fmt.Println(arr)
}
