package main

import (
	"fmt"
	"strconv"
	"strings"
)

// Функция для нахождения общих делителей для массива чисел
func findCommonDivisors(numbers []int) []int {
	if len(numbers) <= 1 {
		return nil
	}
	fmt.Println(len(numbers))
	// Находим наименьшее число в массиве
	minNumber := numbers[0]
	for _, num := range numbers {
		if num < minNumber {
			minNumber = num
		}
	}

	// Инициализируем слайс для хранения общих делителей
	commonDivisors := make([]int, 0)
	fmt.Println("arr:", numbers)
	// Перебираем числа от 1 до минимального числа в массиве
	for divisor := 2; divisor <= minNumber; divisor++ {
		isCommonDivisor := true
		// Проверяем, является ли текущий делитель общим для всех чисел в массиве
		for _, num := range numbers {
			if num%divisor != 0 {
				isCommonDivisor = false
				break
			}
		}
		// Если делитель общий, добавляем его в результат
		if isCommonDivisor {
			commonDivisors = append(commonDivisors, divisor)
		}
	}

	return commonDivisors
}

func main() {
	// Считываем ввод с консоли
	fmt.Println("Введите числа через _ :")
	var input string
	fmt.Scanln(&input)

	// Разбиваем строку на отдельные числа
	numbersStr := strings.Split(input, "_")

	// Преобразуем строковые числа в целые числа
	numbers := make([]int, len(numbersStr))
	for i, numStr := range numbersStr {
		num, err := strconv.Atoi(numStr)
		if err != nil {
			fmt.Println("Ошибка при преобразовании числа:", err)
			return
		}
		numbers[i] = num
	}

	// Находим общие делители
	commonDivisors := findCommonDivisors(numbers)
	fmt.Println("Общие делители:", commonDivisors)
}
