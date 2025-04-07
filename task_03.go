package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 3. Генератор функции для вычисления суммы чисел от 1 до N")
	fmt.Println("=============================================")
	//
	summator := agregator(9) // Получаем функцию-сумматор с ограничением не более 9
	//
	for i := 1; i < 12; i++ {
		fmt.Printf(" = %d\n", summator())
	}
}

// Функция agregator принимает на вход значение-ограничитель
// и возвращает анонимную функцию по суммированию.
func agregator(max int) func() int {
	sum := 0
	current := 0
	return func() int {
		current++
		fmt.Printf("%2d + %2d", sum, current)
		if current <= max {
			sum += current
		}
		return sum
	}
}
