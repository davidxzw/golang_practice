package main

import (
	"fmt"
)

// func PrintType(arr []int){
// 	fmt.Println(reflect.TypeOf(arr).Kind())
// }

func SumFloat64(s []float64) (aver float64) {
	sum := 0.0
	for _, v := range s {
		sum += v
	}
	aver = sum/float64(len(s))
	return 
}

func main() {
	var arr = []float64{1.1, 2.3, 4.5, 100.5}

	aver := SumFloat64(arr)

	fmt.Println(aver)
}
