package main

import (
	"fmt"
	"iter"
	"time"
)

// Let's define an arbitrary struct
type Employee struct {
	Name   string
	Salary int
}

// create a pre-defined list of employees
var Employees = []Employee{
	{Name: "Elliot", Salary: 4},
	{Name: "Donna", Salary: 5},
}

func Iterate[E any](e []E) iter.Seq2[int, E] {
	return func(yield func(int, E) bool) {
		for i := 0; i <= len(e)-1; i++ {
			time.Sleep(1 * time.Second)
			if !yield(i, e[i]) {
				return
			}
		}
	}
}

func IteratePrt() {
	for i, employee := range Iterate(Employees) {
		fmt.Printf("%d: %+v\n", i, employee)
	}
}
