package main

import "fmt"

func main() {
	res := Div(1, 2)

	fmt.Println(res)
}

func Div(a float64, b float64) float64 {
	if b == 0 {
		return 0
	}
	return a / b
}
