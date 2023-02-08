package main

import "fmt"

func test() {
	arr := [5]int{1, 33, 22, 44, 11}

	for i := 0; i < len(arr)-1; i++ {
		for j := 0; j < len(arr)-i-1; j++ {
			if arr[j] > arr[j+1] {
				arr[j], arr[j+1] = arr[j+1], arr[j]
			}
		}
	}

	fmt.Println("arr:", arr)
}
