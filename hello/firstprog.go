package main

import (
	"fmt"
)

func sum(x int, y int) int {
	z := x + y
	return z
}
func main() {
	fmt.Println("hello")

	c := sum
	a := c(1, 2)
	fmt.Println(a)

	//	time.Sleep( time.Minute)

	func(x int) {
		fmt.Println(x)
	}(1)

	func() {
		fmt.Println("hi")
	}()

	//fmt.Println()

}
