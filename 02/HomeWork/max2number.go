package main

import "fmt"

func main() {
	maxnumber := 0
	mang := []int{1, 2, 5, 3}
	for i := 0; i < len(mang); i++ {
		if mang[i] > maxnumber {
			maxnumber = mang[i]
		}
	}
	max2number := 0
	for i := 0; i < len(mang); i++ {
		if mang[i] > max2number && mang[i] < maxnumber {
			max2number = mang[i]
		}
	}
	fmt.Println(max2number)
}
