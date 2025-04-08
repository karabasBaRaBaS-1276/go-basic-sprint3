package main

import (
	"errors"
	"fmt"
)

var ErrTypeTaxi = errors.New("unknown type taxi")
var ErrCalcFare = errors.New("internal calculation error")

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 8. Вычисление стоимости поездки на такси")
	fmt.Println("=============================================")
	//
	lengthKm := 12         // Расстояние в км
	timeTravelMinute := 48 // Время в минутах
	typeTaxi := "business" // Тип такси (econom, standart, business)
	//
	fmt.Printf("Расстояние поездки: %d км\n", lengthKm)
	fmt.Printf("     Время поездки: %d минут\n", timeTravelMinute)
	fmt.Printf("         Тип такси: %s\n", typeTaxi)
	fmt.Println("----------------------------------------")
	//
	coast, cashReceipt, err := calculateFare(typeTaxi, lengthKm, timeTravelMinute)

	if err != nil {
		fmt.Println(errors.Join(errors.New("Ошибка в расчетах:"), err))
		return
	}

	for element, price := range cashReceipt {
		fmt.Printf("%s %4d руб.\n", element, price)
	}

	fmt.Println("----------------------------------------")
	fmt.Printf("Итого к оплате: %d руб.", coast)
}

// Функция получения тарифа за посадку
func getTariffBoard() int {
	return 50
}

// Функция получения тарифа за расстояние по переданному типу такси
func getTariffDistanceByTypeTaxi(typeTaxi string) (int, error) {
	switch typeTaxi {
	case "econom":
		return 10, nil
	case "standart":
		return 15, nil
	case "business":
		return 25, nil
	default:
		return 0, ErrTypeTaxi
	}
}

// Функция получения тарифа за каждую минуту в пути
func getTariffTime() int {
	return 5
}

// Функция рассчета стоимости тарифа
func calculateFare(typeTaxi string, lengthKm int, timeTravelMinute int) (int, map[string]int, error) {
	var (
		cashReceipt map[string]int
		err         error
		coast       int
	)
	cashReceipt = make(map[string]int)
	//
	typeOper := "Оплата посадки в такси:"
	typeOper = fmt.Sprintf("%*s%s",
		(80-len(typeOper))/2,
		" ",
		typeOper,
	)
	sum := getTariffBoard()
	cashReceipt[typeOper] = sum
	coast += sum
	//
	sum, err = getTariffDistanceByTypeTaxi(typeTaxi)
	if err != nil {
		return 0, cashReceipt, fmt.Errorf("%w: %w", ErrCalcFare, err)
	}
	typeOper = fmt.Sprintf("Оплата пути (%d руб. * %d км):", sum, lengthKm)
	typeOper = fmt.Sprintf("%*s%s",
		(70-len(typeOper))/2,
		" ",
		typeOper,
	)
	sum *= lengthKm
	cashReceipt[typeOper] = sum
	coast += sum
	//
	sum = getTariffTime()
	typeOper = fmt.Sprintf("Оплата времени поездки (%d руб. * %d км):", sum, timeTravelMinute)
	typeOper = fmt.Sprintf("%*s%s",
		(70-len(typeOper))/2,
		" ",
		typeOper,
	)
	sum *= timeTravelMinute
	cashReceipt[typeOper] = sum
	coast += sum
	//
	return coast, cashReceipt, nil
}
