package main

import (
	"fmt"
	"strings"
)

type person struct {
	firstname   string
	lastname    string
	fav_country []string
}

type vehicle struct {
	doors int
	color string
}
type truck struct {
	veh       vehicle
	fourWheel bool
}
type sedan struct {
	veh    vehicle
	luxury bool
}

type SHAPE interface {
	area() float64
}
type SQUARE struct {
	side float64
}
type CIRCLE struct {
	r float64
}

func (s SQUARE) area() float64 {
	sqarea := s.side * s.side
	return sqarea
}
func (c CIRCLE) area() float64 {
	carea := 3.14 * c.r * c.r
	return carea
}

func INFO(shape SHAPE) {
	fmt.Println(shape.area())
}

func repetition(st string) map[string]int {

	input := strings.Fields(st)
	wc := make(map[string]int)
	for _, word := range input {
		_, matched := wc[word]
		if matched {
			wc[word] += 1
		} else {
			wc[word] = 1
		}
	}
	return wc
}
func main() {

	val1 := person{
		firstname:   "Alekhya",
		lastname:    "Gandra",
		fav_country: []string{"USA", "Australia", "Italy"},
	}
	val2 := person{
		firstname:   "Akhila",
		lastname:    "xhu",
		fav_country: []string{"Austria", "Turkey", "Dubai"},
	}
	fmt.Println("first person struct", val1)
	fmt.Println("Person First Name ", val1.firstname)
	fmt.Println("Lopping over the favourite Country")
	for i := range val1.fav_country {

		fmt.Println(i, val1.fav_country[i])

	}

	fmt.Println("second person struct", val2)
	fmt.Println("Person second Name ", val2.firstname)
	//fmt.Println(val2.lastname)
	fmt.Println("Lopping over the favourite Country")
	for i := range val2.fav_country {

		fmt.Println(i, val2.fav_country[i])

	}

	mp := map[string]person{

		val1.lastname: val1,
		val2.lastname: val2,
	}
	fmt.Println(mp)
	for _, v := range mp {
		fmt.Println(v)
	}

	truck1 := truck{

		veh:       vehicle{2, "black"},
		fourWheel: true,
	}

	sedan1 := sedan{

		veh:    vehicle{4, "red"},
		luxury: true,
	}

	fmt.Println(truck1)
	fmt.Println(sedan1)
	fmt.Println(truck1.veh)
	fmt.Println(sedan1.luxury)

	s := SQUARE{side: 10}
	c := CIRCLE{r: 5}

	INFO(s)
	INFO(c)

	input := "sun rises in the east"
	for index, element := range repetition(input) {
		fmt.Println(index, "=", element)

	}

}
