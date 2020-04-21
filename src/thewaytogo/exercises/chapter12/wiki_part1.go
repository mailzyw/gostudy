package main

import (
	"flag"
	"fmt"
)

func main()  {

	flag.Parse()
	n := flag.NArg()
	fmt.Println(n)
	fmt.Println(flag.Arg(1))

}



