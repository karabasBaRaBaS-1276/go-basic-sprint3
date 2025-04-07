package main

import "fmt"

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 4. Генератор функции для создания множителей")
	fmt.Println("=============================================")
	//
	multi5 := multiplier(5)
	multi3 := multiplier(3)
	multi2 := multiplier(2)
	//
	for i := 1; i <= 10; i++ {
		fmt.Println("Множитель:", i)
		fmt.Printf("2 х %2d = %2d\n", i, multi2(i))
		fmt.Printf("3 х %2d = %2d\n", i, multi3(i))
		fmt.Printf("5 х %2d = %2d\n", i, multi5(i))
		fmt.Println("--------------")
	}

}

func multiplier(firstMulty int) func(multiplier int) int {
	// firstMulty замыкается, т.к. используется ниже в возвращаемой функции
	return func(secondMulty int) int {
		return firstMulty * secondMulty
	}
}
