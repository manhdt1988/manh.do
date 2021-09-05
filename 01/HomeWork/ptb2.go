package main

import (
	"fmt"
	"math"
)

func main() {

	var a, b, c, x1, x2, Delta float64

	fmt.Print("Enter the a, b, c of Quadratic equation 2 = ")
	fmt.Scanln(&a, &b, &c)

	Delta = (b * b) - (4 * a * c)
	switch {
	case Delta > 0:
		x1 = (-b + math.Sqrt(Delta)/(2*a))
		x2 = (-b - math.Sqrt(Delta)/(2*a))
		fmt.Println("x1 = ", x1, " and x2 = ", x2)
	case Delta == 0:
		x1 = -b / (2 * a)
		x2 = -b / (2 * a)
		fmt.Println("x1 = ", x1, " and x2 = ", x2)
	case Delta < 0:
		x1 := complex(-b/(2*a), math.Sqrt(-Delta)/(2*a))
		x2 := complex(-b/(2*a), -math.Sqrt(-Delta)/(2*a))
		fmt.Println("x1 = ", x1, " and x2 = ", x2)
	default:
		fmt.Println("Please enter Valid Numbers for a,b, and c")
	}
}
