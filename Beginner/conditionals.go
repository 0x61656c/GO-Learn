package main

import "fmt"

func conditionals(){
	i := 0
	for i<= 10{
		if i%2 == 0{
			fmt.Println(true)
		} else {fmt.Println(false)}
		i += 1
	}
}

func main() {
	conditionals()
}
