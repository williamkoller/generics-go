package main

import "fmt"

func indexOf[T comparable](arr []T, s T) int {
	for i, v := range arr {
		if s == v {
			return i
		}
	}

	return -1
}

func main() {
	n := indexOf([]int64{1, 2, 3, 5, 8, 13}, 8)
	fmt.Println("index of int64", n)

	s := indexOf([]string{"py", "js", "go", "rs"}, "js")
	fmt.Println("index of string", s)
}