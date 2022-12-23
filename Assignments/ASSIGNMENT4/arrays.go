package main

import (
	"fmt"
)

func main() {

	/* 1.Create an ARRAY which holds 5 VALUES of TYPE int
	● assign VALUES to each index position.
	● Using format printing ○ print out the TYPE of the array */

	var arr1 [5]int
	arr1[0] = 2
	arr1[1] = 4
	arr1[2] = 6
	arr1[3] = 8
	arr1[4] = 10
	fmt.Printf("The type of the array1 is %T\n", arr1)
	//fmt.Println("Type is: ", reflect.ValueOf(arr1).Kind())

	/* 2.Create a SLICE of TYPE int
		● assign 10 VALUES
	    ● Using format printing ○ print out the TYPE of the slice */

	slice1 := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	fmt.Printf("The type of the array1 is %T\n", slice1)
	//fmt.Println("Type is: ", reflect.ValueOf(slice1).Kind())

	/* 3.use SLICING to create the following new slices which are then printed:
	● [42 43 44 45 46]
	● [47 48 49 50 51]
	● [44 45 46 47 48]
	● [43 44 45 46 47] */

	slice2 := slice1[:5]
	slice3 := slice1[5:]
	slice4 := slice1[2:7]
	slice5 := slice1[1:6]
	fmt.Println("slice2: ", slice2)
	fmt.Println("slice3: ", slice3)
	fmt.Println("slice4: ", slice4)
	fmt.Println("slice5: ", slice5)

	/* 4. Start with this slice
		  ○ x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	      .● append to that slice this value
		  ○ 52
	      ● print out the slice
	          in ONE STATEMENT append to that slice these values
		 ○ 53
		 ○ 54
		 ○ 55
	      ● print out the slice
	      ● append to the x slice the below slice
		○ y := []int{56, 57, 58, 59, 60}
		 ● print out the slice */
	x := []int{42, 43, 44, 45, 46, 47, 48, 49, 50, 51}
	x = append(x, 52)
	fmt.Println("slice x after appending 52: ", x)
	x = append(x, 53, 54, 55)
	fmt.Println("slice x after appending 53,54,55 in one line: ", x)
	y := []int{56, 57, 58, 59, 60}
	x = append(x, y...)
	fmt.Println("Appending y slice to x slice: ", x)

	/* 5.Write a program to pass a pointer to an array as an argument to the function
	   a. Create an array of size 4, data type string
	   b. Create a function with name updateThirdElement and update the value of 3rd index to Texas
	   c. As an input to the function updateThirdElement pass pointer to an array to function updatearray */
	var arr2 = [4]string{"Georgia", "Atlanta", "Chicago", "Dallas"}
	updateThirdElement(&arr2)
	fmt.Println(arr2)

}

func updateThirdElement(p *[4]string) {

	(*p)[3] = "Texas"

}
