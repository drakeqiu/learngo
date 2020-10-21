package main

import "fmt"

func printSlice(s []int) {
	fmt.Printf("%v, len=%d, cap=%d\n",
		s, len(s), cap(s))
}

func deletingElementsFromSlice() {
	fmt.Println("Deleting elements from slice")
	s := []int{2, 4, 6, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	s = append(s[:3], s[4:]...)
	printSlice(s)
}

func popingSlice() {
	fmt.Println("Poping from front")
	s := []int{2, 4, 6, 8, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}
	front := s[0]
	s = s[1:]

	fmt.Println(front)
	printSlice(s)

	fmt.Println("Poping from tail")
	tail := s[len(s)-1]
	s = s[:len(s)-1]

	fmt.Println(tail)
	printSlice(s)
}

func copyingSlice() {
	fmt.Println("Copying slice")
	s1 := []int{2, 4, 6, 8}
	s2 := make([]int, 16)
	copy(s2, s1)
	printSlice(s2)
}

func creatingSlice() {
	fmt.Println("Creating slice")
	var s []int // Zero value for slice is nil
	for i := 0; i < 100; i++ {
		printSlice(s)
		s = append(s, 2*i+1)
	}
	fmt.Println(s)

	s1 := []int{2, 4, 6, 8}
	printSlice(s1)

	s2 := make([]int, 16)
	printSlice(s2)
	s3 := make([]int, 10, 32)
	printSlice(s3)
}

func main() {
	creatingSlice()
	copyingSlice()
	deletingElementsFromSlice()
	popingSlice()
}
