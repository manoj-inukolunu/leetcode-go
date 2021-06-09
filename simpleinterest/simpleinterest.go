package main

import "fmt"

type SalaryCalculator interface {
	CalculateSalary() int
}

type Permanent struct {
	empId    int
	basicpay int
	pf       int
}

type Contract struct {
	empId    int
	basicpay int
}

func (p Permanent) CalculateSalary() int {
	return p.basicpay + p.pf
}

func (c Contract) CalculateSalary() int {
	return c.basicpay
}
func totalExpense(s []SalaryCalculator) {
	expense := 0

	for _, v := range s {
		expense += v.CalculateSalary()
	}
	fmt.Printf("Total Expense Per Month $%d", expense)
}

func main() {
	employees := []SalaryCalculator{Permanent{1, 5000, 20},
		Permanent{2, 6000, 30}, Contract{3, 3000}}
	totalExpense(employees)

}
