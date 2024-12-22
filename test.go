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

func main() {
	var operation Operation

	operation.Name = "молоко"
	operation.Value = 90
	//fmt.Println(operation.Name, operation.Value)

	var category Category
	category.Name = "еда"
	//category.Operations = [

	//fmt.Println(category)

	// создание категории
	var expences Expenses
	var cat string
	var oper string
	var money int
	fmt.Println("Enter category")
	fmt.Scan(&cat)
	expences.Categories.Name = cat
	fmt.Println(expences)

	fmt.Println("Enter operation")
	fmt.Scan(&oper)
	fmt.Println("Enter sum")
	fmt.Scan(&money)

	expences.Categories.Operations.Name = oper
	expences.Categories.Operations.Value = money
	fmt.Println(expences)

}
