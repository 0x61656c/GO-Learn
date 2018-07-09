package main

import "fmt"

func constants(a int) int {
	const x = 1
	return x + a
}


func main() {
	result := constants(1)
}