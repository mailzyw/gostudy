//packege comments
package main

import (
	"fmt"
	"github.com/mailzyw/gostudy/src/stringutil"
)

//the main func
func main() {
	fmt.Println(stringutil.Reverse("!oG,olleH"))
	fmt.Println(Func1(100))
}

//func1 to print num
func Func1(i int) int {
	fmt.Println("num is %d", i)
	return i + 100
}

//func FunctionName(a typea, b typeb)(t1 type1, t2 type2)
//type IZ int
