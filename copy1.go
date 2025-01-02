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
	var choice string

	for outerKey, innerMap := range expences {
		fmt.Printf("%s:\n", outerKey)
		for innerKey, value := range innerMap {
			fmt.Printf("%s = %d\n", innerKey, value)
		}
		fmt.Println()
	}
	fmt.Println("0)Выход в главное меню")
	fmt.Scan(&choice)
	if choice == "0" {
		mainInterface()
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

func mainInterface() {
	fmt.Println("Добро пожаловать в Pig Keeper! Что вы желаете сделать?\n1)Добавить операцию 2)Посмотреть расходы 0)Выход")
	var choice string
	expences := make(map[string]map[string]int)

	expences["еда"] = make(map[string]int)
	expences["еда"]["бананы"] = 100
	expences["еда"]["яблоки"] = 50
	expences["еда"]["молоко"] = 90

	expences["машина"] = make(map[string]int)
	expences["машина"]["бензин"] = 1000
	expences["машина"]["ремонт"] = 2000

	for {
		fmt.Scan(&choice)
		if choice == "1" {
			addOperation(expences)
			fmt.Println("Ваша операция успешно добавлена!\n1)Добавить операцию 2)Посмотреть расходы 0)Выход в главное меню")
			fmt.Scan(&choice)
			if choice == "1" {
				addOperation(expences)
			} else if choice == "2" {
				showData(expences)
			} else if choice == "0" {
				mainInterface()
			}
		} else if choice == "2" {
			showData(expences)
		} else if choice == "0" {
			break
		}
	}
}

func main() {
	mainInterface()
}
