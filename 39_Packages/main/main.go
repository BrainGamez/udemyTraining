package main

import (
	"fmt"

	"github.com/BrainGamez/udemyTraining/01_HelloWorld/hello"
)

func main() {
	fmt.Println(hello.A)
	hello.Change()
	fmt.Println(hello.A)
}
