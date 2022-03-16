package main

import (
	"fmt"

	"golang.org/x/exp/constraints"

	"github.com/go-asset/generics/list"
	"github.com/go-asset/generics/maybe"
)

func add[T constraints.Ordered](c maybe.Maybe[T], value maybe.Maybe[T]) maybe.Maybe[T] {
	if c.IsNothing() {
		return value
	}

	if value.IsNothing() {
		return c
	}

	return maybe.Just(c.Value() + value.Value())
}

func main() {
	result := list.FoldlIter(
		maybe.Just(0),
		readFromUser(),
		add[int],
	)

	resultStr := list.FoldlIter(
		maybe.Just(""),
		readFromUserStr(),
		add[string],
	)

	fmt.Println(result)
	fmt.Println(resultStr)
}
