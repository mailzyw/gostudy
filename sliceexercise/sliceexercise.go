package main

import "fmt"

func main() {
	s := make([][]int, 10)
	s[0][0] = 1
	s[0][5] = 6
	s[9][0] = 91
	s[9][5] = 96
	fmt.Printf("%v\n", s)
}
