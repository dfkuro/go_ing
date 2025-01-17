package main

import "fmt"

func main() {
  var a = "initial"
  fmt.Println(a)

  var b, c int = 1, 2
  fmt.Println(b, c)

  var d = true
  fmt.Println(d)

  var e int
  fmt.Println(e)

  // Declares a new variable "f" with the value "apple" (declares "f" if has not been declared yet)
  f := "apple"
  fmt.Println(f)
}
