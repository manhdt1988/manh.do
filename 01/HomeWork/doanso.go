package main

import (
	"fmt"
)

func main() {
	var i int
	fmt.Println("i =")
	fmt.Scanln(&i)
	if i < 0 {
		fmt.Println("Số bạn đoán nhỏ hơn X")
	} else if (i >= 0) && (i <= 100) {
		fmt.Println("Bạn đã đoán đúng")
	} else if i > 100 {
		fmt.Println("Số bạn đoán lớn hơn X")
	} else {
		fmt.Println("i = ")
	}
}
