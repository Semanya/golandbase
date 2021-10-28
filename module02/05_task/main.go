package main

import "fmt"

func main() {
	check := contains([]string{"lol", "kek", "cheburek"}, "kek")
	fmt.Println(check)
	findMax := getMax(3, 55, 44, 6, 7)
	fmt.Println("Максимальное значение: ", findMax, "     Зуб даю!")
}

func contains(a []string, x string) bool {
	for _, n := range a {
		if x == n {
			return true
		}
	}
	return false
}

func getMax(w ...int) int {
	max := w[0]
	for _, v := range w[1:] {
		if v > max {
			max = v
		}
	}
	return max
}
