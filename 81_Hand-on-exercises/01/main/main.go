package main

// Write a function which takes an integer. The function will have two returns. The first return should be the argument divided by 2. The second return should be a bool that let’s us know whether or not the argument was even. For example:
// half(1) returns (0, false)
// half(2) returns (1, true)


import "fmt"

func main() {
  x := 4
  exone(x)
  fmt.Println(x)
}
// 1. argument delen door twee
// 2. argument
func exone(a int) int, bool {
  z := a/2
  if a%2 == 0 {
    return 1
  }
}
