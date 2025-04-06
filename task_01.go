package main

import (
	"fmt"
	"sort"
	"strings"
)

func main() {
	fmt.Println("=============================================")
	fmt.Println("Задача 1. Подсчет повторяющихся слов в тексте")
	fmt.Println("=============================================")
	//
	var text string = "hello world hello hello everyone"
	fmt.Println("Текст для анализа:\n", text)
	fmt.Println("=============================================")
	//
	fmt.Println("Результат:")
	fmt.Println()
	result, sortWorlds := CalcWorld(text)
	for _, world := range sortWorlds {
		fmt.Printf("%s: %d\n", world, result[world])
	}
	fmt.Println("=============================================")
}

// Подсчет количества слов в тексте без учета регистра.
// Вторым параметром возвращается сипсок слов отсортированных по частоте от большего к меньшему
func CalcWorld(text string) (result map[string]int, sortWorlds []string) {
	result = make(map[string]int)
	arrayWorlds := strings.Fields(strings.ToLower(text))
	for _, world := range arrayWorlds {
		result[world]++
	}
	sortWorlds = make([]string, 0, len(arrayWorlds))
	for wold := range result {
		sortWorlds = append(sortWorlds, wold)
	}

	sort.Slice(sortWorlds, func(i, j int) bool {
		return result[sortWorlds[i]] > result[sortWorlds[j]]
	})
	return
}
