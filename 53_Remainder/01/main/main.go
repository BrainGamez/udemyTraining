package main

import "fmt"

func main() {
	x := 13 / 3
	z := 13 % 3
	fmt.Println("13 / 3 = ", x, " with a remainder of ", z)

	if z == 1 {
		fmt.Println("Which is odd")
	} else {
		fmt.Println("Which is even")
	}

	for i := 1; i < 10; i++ {
		if i%2 == 1 {
			fmt.Println("Odd")
		} else {
			fmt.Println("Even")
		}
	}
}
