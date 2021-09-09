package main

import "fmt"

func removeDuplicate(a []string) (result []string) {
	keys := map[string]bool{}
	for _, entry := range a {
		if _, value := keys[entry]; !value {
			keys[entry] = true
			result = append(result, entry)
		}
	}
	return
}
func main() {
	fmt.Println("-----")
	b := []string{"a", "b", "b", "c", "d", "e", "f", "a"}
	c := removeDuplicate(b)
	fmt.Println(c)
}
