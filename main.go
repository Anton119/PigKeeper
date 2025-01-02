package main

import "fmt"

type Operation struct {
	Name  string
	Value int
}

type Category struct {
	Name       string
	Operations Operation
}

type Expenses struct {
	Categories Category
}

func mainInterface() {
	fmt.Println("Добро пожаловать в Pig Keeper! Что вы желаете сделать?\n1)Добавить операцию 2)Посмотреть расходы 0)Выход")
	var choice string
	for {
		fmt.Scan(&choice)
		if choice == "1" {
			fmt.Println(1)
			break
		} else if choice == "2" {
			fmt.Println(2)
			break
		} else if choice == "0" {
			fmt.Println(0)
			break
		}
	}
}

func main() {
	mainInterface()
}
