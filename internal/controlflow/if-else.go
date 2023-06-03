package controlflow

import "fmt"

func RunIfElse() {
	// if, else
	// ==, !=, >, <, >=, <=
	a := 5
	if a > 100 {
		a = a - 3
	} else if a == 3 {
		a = 0
	} else if a >= 5 {
		a = 100
	} else if a != 8 {
		a = 8
	} else {
		a = a + 3
	}

	fmt.Println(a)
}
