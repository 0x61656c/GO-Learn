package main

import "fmt"

func switchStatements(){
	i := 0
	for i <= 5{ 
		i += 1
		switch i{
		case 1:
			fmt.Println("Case = 1")
		case 2:
			fmt.Println("Case = 2")
		case 3:
			fmt.Println("Case = 3")
		default:
			fmt.Println("This case prints when i >= 4")
		}
	}
}

func main() {
	switchStatements()
}
