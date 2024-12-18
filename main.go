package main

import "fmt"

// считывает строку с терминала
func getInputStr() string {
	var input string
	fmt.Scan(&input)
	return input
}

// добавляет категорию в мапу
func addCategory(m map[string]map[string]int) map[string]map[string]int {
	category := getInputStr()
	m[category] = make(map[string]int)
	return m
}

// добавление операции в мапу
func addOperation(m map[string]map[string]int) map[string]map[string]int {
	fmt.Println("Введите название новой категории: ")
	category := getInputStr()
	fmt.Println("Введите операцию: ")
	operation := getInputStr()
	fmt.Println("Введите сумму операции: ")
	money := getInputInt()

	m[category] = make(map[string]int)
	m[category][operation] = money
	return m
}

// считывает сумму денег с терминала
func getInputInt() int {
	var money int
	fmt.Scan(&money)
	return money
}

// показывает информацию по всей ветеке (расходы/доходы)
func showData(expences map[string]map[string]int) {

	for outerKey, innerMap := range expences {
		fmt.Printf("%s:\n", outerKey)
		for innerKey, value := range innerMap {
			fmt.Printf("%s = %d\n", innerKey, value)
		}
		fmt.Println()
	}
}

// расчитывает общую сумму всех категорий
func countTotalSum(m map[string]map[string]int) int {
	sum := 0
	for _, innerMap := range m {
		for _, value := range innerMap {
			sum += value
		}
	}
	return sum
}

// расчитывает сумму одной категории
func countCategorySum(m map[string]int) int {
	sum := 0
	for _, money := range m {
		sum += money
	}
	return sum
}

func main() {
	expences := make(map[string]map[string]int)

	expences["еда"] = make(map[string]int)
	expences["еда"]["бананы"] = 100
	expences["еда"]["яблоки"] = 50
	expences["еда"]["молоко"] = 90

	expences["машина"] = make(map[string]int)
	expences["машина"]["бензин"] = 1000
	expences["машина"]["ремонт"] = 2000

	showData(expences)
	addOperation(expences)
	showData(expences)
}
