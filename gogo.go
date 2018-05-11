package main

import "fmt"

func main() {
	fmt.Println(fac(5))
}
func fac(n int) int {
	result := 1
	for i := 2; i <= n; i++ {
		result *= i
	}
	return result
}

