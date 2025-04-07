package main

import (
	"fmt"
	"math"
)

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 2. Расчет параметров прямоугольника")
	fmt.Println("=============================================")
	//
	length := 5
	width := 4
	fmt.Printf("Длина: %d; Ширина: %d\n", length, width)
	fmt.Println("=============================================")
	//
	area, perimeter, diagonal := getRectangleData(width, length)
	fmt.Println("Результат:")
	fmt.Println()
	fmt.Printf("Площадь: %d\nПериметр: %d\nДиагональ: %f\n", area, perimeter, diagonal)
	fmt.Println("=============================================")

}

func getRectangleData(width, height int) (area, perimeter int, diagonal float64) {
	area = width * height
	perimeter = (width * 2) + (height * 2)
	diagonal = math.Sqrt(float64(width*width + height*height))
	return
}
