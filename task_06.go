package main

import (
	"errors"
	"fmt"
	"math"
)

var ErrNegative = errors.New("Factorial of a negative number is not defined")
var ErrOverInt = errors.New("Overflow Int type")

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 6. Рекурсивная функция для расчета факториала числа")
	fmt.Println("=============================================")
	//
	a := 5
	//
	result, err := recursiveFactorial(a)
	if err != nil {
		fmt.Printf("Ошибка при расчете факториала числа %d: %s", a, err)
		return
	}
	fmt.Printf("%d! = %d", a, result)
}

func recursiveFactorial(n int) (int, error) {
	if n < 0 {
		return 0, ErrNegative
	}
	if (n == 1) || (n == 0) {
		return 1, nil
	}
	prev, err := recursiveFactorial(n - 1) // Факториал предыдущего числа
	if err != nil {
		return 0, err
	}

	if prev > math.MaxInt/n {
		return 0, ErrOverInt
	}
	return n * prev, nil
}
