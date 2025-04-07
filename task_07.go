package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 7. Анонимная функция для вычисления суммы чисел в слайсе")
	fmt.Println("=============================================")
	//
	sliceInt := []int{1, 5, 8, 2, 9, 0, -5}
	fmt.Println("Дано:", sliceInt)
	//
	sum := func(slice []int) (result int) {
		result = 0
		for _, value := range slice {
			result += value
		}
		return
	}
	// Вызов функции
	result := sum(sliceInt)
	fmt.Println("Сумма:", result)
}
