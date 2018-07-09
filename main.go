package main

import "fmt"

func helloWorld(a string) string {
	return a
}

func constants(a int) int {
	const x = 1
	return x + a
}

func variables(a int, b int) int { 
	return a + b
}

func loops(a int) int{
	i := 0
	sum := 0

	for i <= a{
		sum = sum + i
		i = i + 1
	}
	return sum
}

func conditionals(){
	i := 0
	for i<= 10{
		if i%2 == 0{
			fmt.Println(true)
		} else {fmt.Println(false)}
		i += 1
	}
}

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

func arrays(){
	
	var a [5]int
	fmt.Println(a)

	i := 0
	for i < 5{
		a[i] = i
		i = i + 1
	}
	fmt.Println(a)
}


func main() {
	a := helloWorld("1")
	b := constants(1)
	c := variables(1,2)
	d := loops(4) 

	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(c)
	fmt.Println(d)

	conditionals()
	switchStatements()
	arrays()

}

