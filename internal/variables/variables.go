package variables

import "fmt"

var globVar = "GG"

func RunVariables() {
	hello := "World"
	fmt.Println(hello)

	var helloVar string
	helloVar = "Ake"
	fmt.Println(helloVar)

	var goodVar = "Earth"
	fmt.Println(goodVar)

	fmt.Println(globVar)
}
