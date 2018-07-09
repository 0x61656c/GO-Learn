package main

import "fmt"

func loops(a int) int{
	i := 0
	sum := 0

	for i <= a{
		sum = sum + i
		i = i + 1
	}
	return sum
}
func main() {
	result := loops(4)
}
