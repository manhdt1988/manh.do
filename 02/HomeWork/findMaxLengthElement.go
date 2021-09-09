package main

import "fmt"

func main() {
	mang := []string{"aa", "Tuan", "bcd", "a", "Manh", "bb"}
	maxLengthElement := 0
	ketqua := []string{}
	for i := 0; i < len(mang); i++ {
		lengthElement := len(mang[i])
		if lengthElement > maxLengthElement {
			maxLengthElement = lengthElement
		}
	}
	for _, value := range mang {
		if len(value) == maxLengthElement {
			ketqua = append(ketqua, value)
		}
	}
	fmt.Println(ketqua)
}
