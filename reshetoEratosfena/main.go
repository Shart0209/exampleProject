package main

import "fmt"

func generateSlice(x int) []int {
	var sl []int
	for i := 2; i < x; i++ {
		sl = append(sl, i)
	}
	return sl
}

func removeIndex(s *[]int, index int) {
	*s = append((*s)[:index], (*s)[index+1:]...)
}

func main() {
	n := 121
	items := generateSlice(n)

	for i, p := range items {
		if p*p > n {
			break
		}
		for j := i + 1; j < len(items); j++ {
			if items[j]%p == 0 {
				removeIndex(&items, j)
			}
		}
	}

	fmt.Printf("Для n = %d\nРезультат = %v", n-1, items)

}
