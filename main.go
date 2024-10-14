package main

import "fmt"

func add(a int, b int) int {
	return a + b
}
func minus(a int, b int) int {
	return a - b
}
func summ(k int) int {
	sum := 0
	for i := 0; i < k; i++ {
		sum += i
		fmt.Print(" ", sum)
	}
	return sum
}
func factorial(n int) int {
	res := 1
	for i := 1; i <= n; i++ {
		res *= i
		fmt.Println(i, "! = ", res)
	}
	return res
}
func main() {
	fmt.Println("Сумма равна ", add(10, 19))
	fmt.Println("Разность равна ", minus(100, 19))
	fmt.Println(" нарастающий итог ", summ(5))
	fmt.Println("5! = ", factorial(5))
}
