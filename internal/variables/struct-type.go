package variables

import "fmt"

type Person struct{
  Name    string
  Age     int
  IsAdmin bool
}

func RunStructType() {
  p := Person{
    Name:    "Mike",
    Age:     16,
    IsAdmin: false,
  }
  fmt.Println(p, p.Name, p.Age, p.IsAdmin)
}
