package main

import "fmt"

func main() {
	var num, i int
	var arr [10]int
	fmt.Println("Enter the number you want to convert:")
	fmt.Scan(&num)
	for i = 0; num > 0; i++ {
		arr[i] = num % 2

		num = num / 2
		if num < 1 {
			break
		}
	}
	answer := "The Binary of given number is "
	for i >= 0 {
		answer = fmt.Sprint(answer, arr[i])
		i--
	}
	fmt.Println(answer)
}
