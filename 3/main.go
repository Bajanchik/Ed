package main

import (
	"fmt"
)

// Функция для определения является ли число простым
func isPrime(num int) bool {
	if num <= 1 {
		return false
	}
	if num <= 3 {
		return true
	}
	if num%2 == 0 || num%3 == 0 {
		return false
	}
	i := 5
	for i*i <= num {
		if num%i == 0 || num%(i+2) == 0 {
			return false
		}
		i += 6
	}
	return true
}

// Функция для нахождения простых чисел в заданном диапазоне
func findPrimesInRange(min, max int) []int {
	primes := make([]int, 0)
	for num := min; num <= max; num++ {
		if isPrime(num) {
			primes = append(primes, num)
		}
	}
	return primes
}

func main() {
	// Указываем диапазон
	fmt.Println("Введите начало диапазона: ")
	var min int
	_, err1 := fmt.Scan(&min)
	if err1 != nil {
		fmt.Printf("Ошибка сканирования: %v", err1)
	}
	fmt.Println("Введите конец диапазона: ")
	var max int
	_, err2 := fmt.Scan(&max)
	if err2 != nil {
		fmt.Printf("Ошибка сканирования: %v", err2)
	}

	// Находим простые числа в заданном диапазоне
	primes := findPrimesInRange(min, max)

	// Выводим результат
	fmt.Printf("Простые числа в диапазоне от %d до %d: %v\n", min, max, primes)
}
