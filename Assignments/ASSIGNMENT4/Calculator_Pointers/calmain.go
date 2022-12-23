package main

import (
	"fmt"
	"golang/Arithmetic"
)

func main() {

	/* var a, b int
	fmt.Println("Enter a Value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b Value: ")
	fmt.Scanln(&b) */
	a := Arithmetic.X
	b := Arithmetic.Gety()

	fmt.Println(" Addition : ", Arithmetic.Addition(&a, &b))
	fmt.Println(" Subtraction : ", Arithmetic.Subtraction(&a, &b))
	fmt.Println(" Multiplication : ", Arithmetic.Multiplication(&a, &b))
	quo, rem := Arithmetic.Division(&a, &b)

	fmt.Println(" Division : ")
	fmt.Println("   Quotient : ", quo)
	fmt.Println("   Reminder : ", rem)
}
