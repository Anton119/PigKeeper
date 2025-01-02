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

func (c *Category) AddOperationToCategory(operation Operation) Category {
	var category Category
	operations := []Operation{}

	name := "food"
	operations = append(operations, operation)

	c.Name = name
	c.Operations = operations
	return category
}

// / creat method  + add function
func (o *Operation) CreateOperation(operation string, money int) {
	o.Name = operation
	o.Value = money
}

func AddOperation() Operation {
	var operation Operation
	var op string
	var sum int

	/*
		fmt.Println("Enter an operation")
		fmt.Scan(&op)
		fmt.Println("Enter money")
		fmt.Scan(&sum)
	*/

	op = "banana"
	sum = 90

	operation.CreateOperation(op, sum)
	return operation

}

func main() {
	op := AddOperation()
	fmt.Println(op)

	var cat Category
	fmt.Println(cat)
	cat.AddOperationToCategory(op)
	fmt.Println(cat)

	/*var expenses Expenses

	var cat string
	fmt.Println("Enter a category")
	fmt.Scan(&cat)

	expenses.CreateCategory(cat)
	fmt.Println(expenses)
	*/

	/*var operation Operation

	operation.Name = "молоко"
	operation.Value = 90
	//fmt.Println(operation.Name, operation.Value)

	var category Category
	category.Name = "еда"
	fmt.Println(category)

	var expenses Expenses
	expenses.Name = "расходы"
	fmt.Println(expenses.Categories)

	expenses.addCategory()
	fmt.Println(expenses.Categories)

	*/
	//category.Operations = [

	//fmt.Println(category)

	// создание категории

	/*var expences Expenses
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
	*/

}
