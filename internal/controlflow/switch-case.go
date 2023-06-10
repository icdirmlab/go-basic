package controlflow

import "fmt"

func RunSwitchCase() {
	b := "NO"
  switch b {
  case "YES":
    b = "Y"
  case "NO":
    b = "N"
  default:
    b = "X"
  }
  fmt.Println(b)
}
