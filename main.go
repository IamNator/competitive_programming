package main

import "fmt"

func main() {
	testSlice := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Printf("sample %v\n", testSlice)
	ans := addByPtr(&testSlice, 1, 1, 1)
	fmt.Printf("input %v, output %v\n", testSlice, ans)

}

func sq(s []int) {
	for i, v := range s {
		s[i] = v * v
	}
}

func addCpy(s []int, n ...int) []int {
	return append(s, n...)
}

func addByPtr(s *[]int, n ...int) []int {
	*s = append(*s, n...)

	return *s
}
