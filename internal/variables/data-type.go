package variables

import "fmt"

func RunDataType() {
	a := 1    // var a int
	b := 3.14 // var b float
	c := "hi" // var c string
	d := true // var d bool

	fmt.Println(a, b, c, d)

	e := []int{1, 2, 3} // slice
	e = append(e, 4)
	// [start index:cout end]
	fmt.Println(e, len(e), e[0], e[1:3], e[1:], e[:3])

	f := make(map[string]int) // map [key]value
	f["one"] = 1
	f["two"] = 2
	fmt.Println(f, len(f), f["one"], f["three"])
}
