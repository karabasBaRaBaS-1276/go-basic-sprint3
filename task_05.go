package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 5. Передача функций в операцию")
	fmt.Println("=============================================")
	//
	a := 15
	b := 3
	//
	fmt.Printf("Дано:\na = %3d\nb = %3d\n", a, b)
	fmt.Println("=============================================")
	//
	var mapFunc map[string]func(int, int) int
	mapFunc = make(map[string]func(int, int) int, 10)
	mapFunc["Сумма:        "] = add
	mapFunc["Разность:     "] = subtract
	mapFunc["Произведение: "] = func(a int, b int) int {
		return a * b
	}
	mapFunc["Частное:      "] = func(a int, b int) int {
		return a / b
	}

	//
	for name, operation := range mapFunc {
		fmt.Printf("%s %3d\n", name, applyOperation(a, b, operation))
	}
}

// Операции над числами. Сама операция передается в виде функции operation
func applyOperation(a int, b int, operation func(int, int) int) (result int) {
	result = operation(a, b)
	return
}

func add(a, b int) int {
	return a + b
}

func subtract(a, b int) int {
	return a - b
}
