package Arithmetic

var X int = 10
var y int = 5

func Gety() int {
	return y
}

func Addition(X int, y int) int {
	return X + y
}

func Subtraction(X int, y int) int {
	return X - y
}

func Multiplication(X int, y int) int {
	return X * y
}

func Division(X int, y int) (int, int) {
	return X / y, X % y
}
