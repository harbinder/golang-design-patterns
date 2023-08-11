package sort

import (
	"fmt"
	"sort"
)

type Employee struct {
	name   string
	age    int
	salary float32
}

type Department []Employee

func (d Department) Len() int {
	return len(d)
}

func (d Department) Less(i, j int) bool {
	if d[i].age < d[j].age {
		return true
	} else if d[i].age == d[j].age {
		return d[i].salary < d[j].salary
	} else if d[i].salary == d[j].salary {
		return d[i].name < d[j].name
	}
	return false
}

func (d Department) Swap(i, j int) {
	temp := d[i]
	d[i] = d[j]
	d[j] = temp
}

func ExampleSort() {
	dp := Department{
		{name: "Harry", age: 65, salary: 1000},
		{name: "Shaun", age: 25, salary: 5000},
		{name: "Brown", age: 25, salary: 5000},
		{name: "Glassman", age: 65, salary: 9000},
		{name: "Lea", age: 25, salary: 5000},
	}
	fmt.Println("Pre-Sort", dp)
	sort.Sort(dp)
	fmt.Println("Post-Sort", dp)
}
