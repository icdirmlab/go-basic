package variables

import "fmt"

func RunZeroValues() {
	var a int
  var b float64
  var c bool
  var d string
  fmt.Println(a, b, c, d)
  
  var e []int
  var f map[string]float64
  fmt.Println(e, f)
  
  type person struct{
    name string
    age int
  }
  var g person
  var h *person
  fmt.Println(g, h)
}
