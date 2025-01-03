package main

import "fmt"

type Operation struct {
	Name  string
	Value int
}

type Category struct {
	Name       string
	Operations []Operation
}

type Expenses struct {
	Name       string
	Categories []Category
}

func choiceInput() string {
	var choice string
	fmt.Scan(&choice)
	return choice
}

func (o *Operation) CreateOperation() Operation {
	var operation Operation
	var op string
	var sum int

	fmt.Println("Введите название операции")
	fmt.Scan(&op)
	fmt.Println("Введите сумму операции")
	fmt.Scan(&sum)

	o.Name = op
	o.Value = sum
	return operation

}

func (c *Category) CreateCategory() Category {
	var category Category
	var name string

	fmt.Println("Введите имя категории")
	fmt.Scan(&name)
	c.Name = name
	return category
}

// запрашивает имя из терминала
func getName() string {
	var name string
	fmt.Scan(&name)
	return name
}

func (c *Category) AddOperationToCategory(name string, operation Operation) *Category {

	c.Name = name
	c.Operations = append(c.Operations, operation)
	return c
}

func mainInterface() {
	fmt.Println("Добро пожаловать в Pig Keeper! Что вы желаете сделать?\n1)Добавить операцию 2)Посмотреть расходы 0)Выход")
	for {
		choice := choiceInput()
		if choice == "1" {
			fmt.Println("1)Выбрать категорию 2)Создать категорию")
			choice = choiceInput()
			if choice == "1" {
				// выбрать категорию
				fmt.Println("1)Еда")
				choice = choiceInput()
				if choice == "1" {
					var operation Operation
					operation.CreateOperation()
					fmt.Println(operation)
					fmt.Println("1)Добавить операцию 2)Назад 3)Выйти в главное меню")
					choice = choiceInput()
					if choice == "1" {
						var operation Operation
						operation.CreateOperation()
						fmt.Println(operation)
					} else if choice == "2" {
						break
					}
				}
			} else if choice == "2" {
				// создать категорию
				var category Category
				category.CreateCategory()
				fmt.Println(category)

				fmt.Println("1)Создать операцию 2)Назад")
				choice = choiceInput()
				if choice == "1" {
					// создать операцию
					var operation Operation

					operation.CreateOperation()
					category.AddOperationToCategory(category.Name, operation)
					fmt.Println(category)
					continue
				} else if choice == "2" {
					break
				}
				break
			}

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
