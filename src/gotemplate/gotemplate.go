package main

import (
	"fmt"
	"os"
	"runtime"
)

const c = "C"

var v int = 5

type T struct{}

func init() { // initialization of package
}

func main() {
	var a int
	b := 100
	Func1()
	// ...
	fmt.Println(a)

	var goos string = runtime.GOOS
	fmt.Printf("The operating system is %s\n", goos)
	path := os.Getenv("PATH")
	fmt.Printf("Path is %s\n", path)
	fmt.Println("Hello:", 23, 24, "aaa", b)
}

func (t T) Method1() {
	//...
}

func Func1() { // exported function Func1
	//...
}
