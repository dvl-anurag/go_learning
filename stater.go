package main

import (
	"fmt"
	b "goTour/basics"
)

func main() {

	fmt.Println("Started from main")

	fmt.Println(b.WithParam(10, 20))
	fmt.Println(b.Swap(10, 20))
	fmt.Println(b.Split(10))
	b.BaseType()
	b.ZeroVal()
	b.ConverterTool()
	b.Consta()
}
