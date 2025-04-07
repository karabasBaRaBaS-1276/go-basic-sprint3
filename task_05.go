package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 5. Передача функций в операцию")
	fmt.Println("=============================================")
	//
	a := 15
	b := 0
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
		// тут может быть деление на 0. Специально "забыли" про проверку, чтобы проверить работу defer
		return a / b
	}

	//
	for name, operation := range mapFunc {
		fmt.Printf("%s %3d\n", name, applyOperation(a, b, operation))
	}
}

// Операции над числами. Сама операция передается в виде функции operation
func applyOperation(a int, b int, operation func(int, int) int) (result int) {
	// Так как нет гарантии как хорошо написаны функции, то имеет смысл добавить восстановление в случае ошибок
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Восстановление работоспособности после ошибки. ", r)
		}
	}()
	result = operation(a, b)
	return
}

// Сложение
func add(a, b int) int {
	return a + b
}

// Вычитание
func subtract(a, b int) int {
	return a - b
}
