package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// 1. Using the for loop print 1 to 100 numbers
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
	}
	//2. Create a for loop using this syntax for condition { } print out the odd numbers in 1 to 50.
	i := 1
	for i <= 50 {
		if i%2 != 0 {
			fmt.Print(i, " ")
		}
		i++
	}
	fmt.Println()
	// 3. Create a for loop using this syntax for { } print out the even numbers in 1 to 50.
	j := 1
	for {

		if j%2 == 0 {
			fmt.Print(j, " ")

		} else if j > 50 {
			break
		}
		j++
	}
	fmt.Println()
	// 4. Print out the quotient for each number between 50 and 105 when it is divided by 6.
	for i := 50; i <= 105; i++ {
		fmt.Println(i / 6)
	}
	// 5. Read the user input.If the input is equal to Golang tutorial print welcome else print end

	fmt.Print("Enter your String: \n")
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()
	str := scanner.Text()

	if str == "Golang tutorial" {
		fmt.Println("Welcome")
	} else {
		fmt.Println("End")
	}

	//6.Write a program to print the numbers from 1 to 80.
	//But, for multiples of two print Golang instead of the number and for the multiples of four print tutorial.
	// For numbers which are multiples of both two and four print Golang tutorial

	for i := 1; i <= 80; i++ {
		if i%2 == 0 && i%4 == 0 {
			fmt.Println("Golang tutorial")

		} else if i%2 == 0 {
			fmt.Println("Golang")

		} else if i%4 == 0 {
			fmt.Println("tutorial")

		} else {
			fmt.Println(i)
		}

	}

}
