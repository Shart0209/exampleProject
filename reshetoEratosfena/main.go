package main

import (
	"fmt"
)

func generateSlice(x int) []int {
	var sl []int
	for i := 2; i < x; i++ {
		sl = append(sl, i)
	}
	return sl
}

func removeIndex(s []int, index int) []int {
	return append(s[:index], s[index+1:]...)
}

func main() {
	n := 121
	items := generateSlice(n)

	for i, p := range items {
		for j := i + 1; j < len(items); j++ {
			switch {
			case p*p > n:
				break
			case items[j]%p == 0:
				sl := removeIndex(items, j)
				items = sl
			}
		}
	}

	fmt.Printf("Для n = %d\nРезультат = %v", n-1, items)

}
