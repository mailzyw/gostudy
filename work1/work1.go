package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs(a int) float64 {
	return float64(math.Sqrt(v.X*v.X + v.Y*v.Y + float64(a)))
}

func main() {
	v := Vertex{3, 4}
	fmt.Println(v.Abs(10))
}
