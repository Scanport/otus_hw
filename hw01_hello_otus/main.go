package main

import (
	"fmt"

	"golang.org/x/example/stringutil"
)

func main() {
	input := "Hello, OTUS!"
	output := stringutil.Reverse(input)
	fmt.Println(output)
}
