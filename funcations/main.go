package main

import "fmt"

// Multiple results
func swap(x, y string) (string, string) {
	return y, x
}

func add(x, y int) int {
	return x + y
}

// Named return values
func split(sum int) (x, y int) {
	x = sum * 4 / 9
	y = sum - x
	return
}

func main() {
	fmt.Println(add(10, 20))
	fmt.Println(swap("Hello", "world"))
	fmt.Println(split(17))
}
