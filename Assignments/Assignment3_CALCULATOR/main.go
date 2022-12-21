package main

import (
	"fmt"

	"Assignments/Arithmetic"
)

func main() {

	/* var a, b int
	fmt.Println("Enter a Value: ")
	fmt.Scanln(&a)
	fmt.Println("Enter b Value: ")
	fmt.Scanln(&b) */

	fmt.Println(" Addition : ", Arithmetic.Addition(Arithmetic.X, Arithmetic.Gety()))
	fmt.Println(" Subtraction : ", Arithmetic.Subtraction(Arithmetic.X, Arithmetic.Gety()))
	fmt.Println(" Multiplication : ", Arithmetic.Multiplication(Arithmetic.X, Arithmetic.Gety()))
	quo, rem := Arithmetic.Division(Arithmetic.X, Arithmetic.Gety())

	fmt.Println(" Division : ")
	fmt.Println("   Quotient : ", quo)
	fmt.Println("   Reminder : ", rem)
}
