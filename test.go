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

func (ex *Expenses) AddCategoryToExpenses(category Category) *Expenses {

	name := "expences"

	ex.Name = name
	ex.Categories = append(ex.Categories, category)
	return ex

}

/*func (ex *Expenses) listOfCategories() []string {
	categories := []string{}
	categories = append(categories, ex.Categories)
	return categories

}*/

// считает кол-во категорий
func (exp *Expenses) countCategories() int {
	count := 0
	for i := 0; i < len(exp.Categories); i++ {
		count++
	}
	return count
}

// проверяет есть ли категория уже в списке
func (exp *Expenses) categoryExists(count int, spisokCategories []string) bool {

	for i := 0; i < count; i++ {
		if "ff" == "food" {
			return true
		}
	}
	return false
}

// запрашивает имя из терминала
func getName() string {
	var name string
	fmt.Scan(&name)
	return name
}

func (c *Category) AddOperationToCategory(operation Operation) *Category {

	c.Name = getName()
	c.Operations = append(c.Operations, operation)
	return c
}

// / creat method  + add function
/*func (o *Operation) CreateOperation(operation string, money int) {
	o.Name = operation
	o.Value = money
}*/

func (o *Operation) CreateOperation() Operation {
	var operation Operation
	var op string
	var sum int

	fmt.Println("Enter an operation")
	fmt.Scan(&op)
	fmt.Println("Enter money")
	fmt.Scan(&sum)

	o.Name = op
	o.Value = sum
	return operation

}

func (c *Category) CreateCategory() Category {
	var category Category
	var name string
	fmt.Scan(&name)
	c.Name = name
	return category
}

func main() {

	var cat Category
	cat.CreateCategory()
	cat.AddOperationToCategory()
	fmt.Println(cat)

	var ex Expenses

	op1 := Operation{
		Name:  "banana",
		Value: 200,
	}

	cat.AddOperationToCategory(op1)
	fmt.Println(cat)

	ex.AddCategoryToExpenses(cat)

	op2 := Operation{
		Name:  "ticket",
		Value: 300,
	}

	var cat2 Category

	fmt.Println("введи имя категории")
	cat2.AddOperationToCategory(op2)
	fmt.Println(cat2)

	ex.AddCategoryToExpenses(cat2)
	fmt.Println(ex)

	//arr := ex.listOfCategories()
	//fmt.Println(arr)
	/*listOfCat := []string{}
	for i := 0; i < ex.countCategories(); i++ {
	}
	*/
	/*
		var op Operation
		op.CreateOperation()
		fmt.Println(op)

		var cat Category
		fmt.Println(cat)
		cat.AddOperationToCategory(op)
		fmt.Println(cat)

		var exp Expenses
		fmt.Println(exp)
		exp.AddCategoryToExpenses(cat)
		fmt.Println(exp)
		//////////////////////
		fmt.Println("Add 1 more operation")
		op.CreateOperation()
		fmt.Println(op)

		fmt.Println(cat)
		cat.AddOperationToCategory(op)
		fmt.Println(cat)

		fmt.Println(exp)
		exp.AddCategoryToExpenses(cat)
		fmt.Println(exp)

	*/

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
