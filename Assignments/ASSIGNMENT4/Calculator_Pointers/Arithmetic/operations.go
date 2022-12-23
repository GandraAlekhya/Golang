package Arithmetic

var X int = 10
var y int = 5

func Gety() int {
	return y
}

func Addition(p1 *int, p2 *int) int {
	r1 := (*p1) + (*p2)
	return r1
}

func Subtraction(p1 *int, p2 *int) int {
	r1 := (*p1) - (*p2)
	return r1
}

func Multiplication(p1 *int, p2 *int) int {
	r1 := (*p1) * (*p2)
	return r1
}

func Division(p1 *int, p2 *int) (int, int) {
	r1 := (*p1) / (*p2)
	r2 := (*p1) % (*p2)
	return r1, r2
}
