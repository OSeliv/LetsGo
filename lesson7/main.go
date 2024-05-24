package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("Task 7.1")
	var arr1 [5]string // = [5]string{"Mon", "Tue", "Wed", "Thu", "Fri"}
	fmt.Printf("%#v\n", arr1)

	fmt.Println("Task 7.2")
	arr2 := [4]string{"яблоко", "груша", "слива", "абрикос"}
	fmt.Println(arr2)

	fmt.Println("Task 7.3")
	arr3 := [4]string{"яблоко", "груша", "помидор", "абрикос"}
	arr3[2] = "персик" //changing tomato
	fmt.Println(arr3)

	fmt.Println("Task 7.4")
	s4 := []int{5, 2, 8, 3, 1, 9}
	fmt.Println(s4)

	fmt.Println("Task 7.5")
	s5 := make([]int, 0, 10) // 7.5 + 7.6
	fmt.Println(s5, "len:", len(s5), "cap:", cap(s5))

	fmt.Println("Task 7.6")
	// s5 - empty slice with capacity = 10
	s5 = append(s5, 4, 1, 8, 9) // to append elements
	fmt.Println(s5, "len:", len(s5), "cap:", cap(s5))

	fmt.Println("Task 7.7")
	s71 := []int{1, 2, 3}     // 1st slice
	s72 := []int{4, 5, 6}     //second slice
	s7 := append(s71, s72...) // resulting slice
	fmt.Println(s7)

	fmt.Println("Task 7.8")
	s8 := []int{1, 2, 3, 4, 5, 6}
	i := sort.SearchInts(s8, 4)      // searching element
	s8 = append(s8[:i], s8[i+1:]...) //union before element and after
	fmt.Println(s8)
}
