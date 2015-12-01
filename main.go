package main

import "fmt"

func factorial(x uint) uint {
	//fmt.Println("x:", x)
	if x == 0 {
		return 1
	}
	return x * factorial(x-1)
}

func fibonacci(x uint) uint {
	//fmt.Println("x:", x)
	if x == 0 {
		return 0
	} else if x == 1 {
		return 1
	}
	return fibonacci(x-1) + fibonacci(x-2)
}
func main() {
	fmt.Println("FACTORIAL:", factorial(10))
	fmt.Println("FIBONACCI:", fibonacci(10))
}
