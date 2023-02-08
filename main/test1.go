package main

import "fmt"

func test1() {
  arr := [5]int{22, 11, 33, 44, 55}
  
  //选择排序
  for i := 0; i < len(arr)-1; i++ {
    min := i
    for j := i+1; j < len(arr); j++ {
      if arr[j] < arr[min] {
        min = j
      }
    }
    if min != i {
      arr[i], arr[min] = arr[min], arr[i]
    }
  }
  
  fmt.Pritnln("arr:", arr)
}
