package main

import (
	"fmt"
)

func computer(number int) string {
	var word string

	// Определяем падеж в зависимости от числа
	if number%10 == 1 && number%100 != 11 {
		word = "компьютер"
	} else if (number%10 >= 2 && number%10 <= 4) && !(number%100 >= 12 && number%100 <= 14) {
		word = "компьютера"
	} else {
		word = "компьютеров"
	}

	return word
}

func main() {
	var number int
	fmt.Println("Enter your number: ")

	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("error with scan num: %v", err)
	}

	res := computer(number)
	fmt.Println(number, res)
}
