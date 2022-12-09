package main

import (
	"fmt"
	monkey_sort "github.com/golang-infrastructure/go-monkey-sort"
)

func main() {
	slice := []int{2, 3, 1, 4}
	count := monkey_sort.Sort(slice)
	fmt.Println(count)
	fmt.Println(slice)
	// Output:
	// 4861854
	// [1 2 3 4]
}
