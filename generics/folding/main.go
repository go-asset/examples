package main

import (
	"fmt"

	"github.com/go-asset/examples/internal/filter"
	"github.com/go-asset/generics/list"
)

func main() {
	myList := []int{1, 2, 3, 4, 5}

	resultInt := list.Foldl(int64(myList[0]), myList[0:], func(c int64, value int) int64 {
		return c * int64(value)
	})

	resultStr := list.Foldr("", []rune("fancy"), func(c string, value rune) string {
		return c + string(value)
	})

	people := []Person{
		{Name: "Alice", Age: 16},
		{Name: "Bob", Age: 21},
		{Name: "Carol", Age: 33},
		{Name: "Dylan", Age: 12},
	}

	resultFamily := list.Foldl(Family{}, people, func(c Family, value Person) Family {
		if value.Age < 18 {
			c.Children = append(c.Children, value)
		} else {
			c.Adults = append(c.Adults, value)
		}

		return c
	})

	fmt.Println(resultInt)
	fmt.Println(resultStr)
	fmt.Println(resultFamily.Adults)
	fmt.Println(resultFamily.Children)

	query := filter.And(
		filter.Or(filter.Name("alice"), filter.Name("bob")),
		filter.Or(filter.Age(55), filter.Age(31)),
	)

	fmt.Println(query.String())
}
