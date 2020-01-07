package main

import "fmt"

func main() {
	sliceA := []int{1, 2, 3}
	sliceB := []int{4, 5, 6}

	sliceA = append(sliceA, sliceB)
	//sliceA = append(sliceA, 4, 5, 6)

	fmt.Println(sliceA) // [1 2 3 4 5 6] 출력
	fmt.Println(sliceB)
}
