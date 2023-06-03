package main

import (
	"fmt"
	"go-basic/internal/variables"
)

func main() {
	fmt.Println("---- RunVariables -----")
	variables.RunVariables()
	fmt.Println()
	fmt.Println("---- RunDataType -----")
	variables.RunDataType()
	fmt.Println()
	fmt.Println("---- RunStructType -----")
	variables.RunStructType()
	fmt.Println()
	fmt.Println("---- RunTypeConvertion -----")
	variables.RunTypeConvertion()
	fmt.Println()
	fmt.Println("---- RunZeroValues -----")
	variables.RunZeroValues()
}
