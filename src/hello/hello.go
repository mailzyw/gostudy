//packege comments
package main

import (
	"bufio"
	"fmt"
	"os"
)

//the main func
func main() {
	file, err := os.Open("products.txt")
	if err!=nil{
		fmt.Printf("open file error")
		return
	}
	defer file.Close()

	iReader := bufio.NewReader(file)
	for{
		str,err := iReader.ReadString('\n')
		if err != nil{
			return
		}
		fmt.Printf("The input was : %s",str)
	}

}











