package main

import "fmt"

// Метод для вывода таблицы умножения до указанного числа
func multiplicationTable(n int) {
	fmt.Printf("Таблица умножения до %d:\n", n)
	for i := 1; i <= n; i++ {
		for j := 1; j <= n; j++ {
			fmt.Printf("%d\t", i*j)
		}
		fmt.Println()
	}
}

func main() {
	// Указываем число до которого выводим таблицу
	fmt.Println("Введите число до которого строить таблицу умножения: ")
	var number int
	_, err := fmt.Scan(&number)
	if err != nil {
		fmt.Printf("Ошибка сканирования: %v", err)
	}

	// Выводим таблицу умножения
	multiplicationTable(number)
}
