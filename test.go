package main

import (
	"fmt"
)

func main() {
	// Создаем мапу, где ключи - строки, а значения - мапы
	nestedMap := make(map[string]map[string]int)

	// Инициализируем вложенные мапы
	nestedMap["first"] = make(map[string]int)
	nestedMap["first"]["a"] = 1
	nestedMap["first"]["b"] = 2

	nestedMap["second"] = make(map[string]int)
	nestedMap["second"]["c"] = 3
	nestedMap["second"]["d"] = 4

	// Выводим значения
	for outerKey, innerMap := range nestedMap {
		fmt.Printf("%s: ", outerKey)
		for innerKey, value := range innerMap {
			fmt.Printf("%s=%d ", innerKey, value)
		}
		fmt.Println()
	}
}
