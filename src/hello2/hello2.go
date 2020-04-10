package main

import (
	"fmt"
	"hello2/helloutil"
)

func main() {
	helloutil.PrintHello("Hello,world.\n")
	helloutil.PrintHello("Hello,world2.\n")
	i := 10
	j := 100
	fmt.Printf("Hello,world3.%d\n", i)
	fmt.Printf("Hello,world4.%d\n", j)
}
