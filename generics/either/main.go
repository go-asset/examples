package main

import (
	"fmt"
	"log"

	"golang.org/x/exp/constraints"

	"github.com/go-asset/generics/either"
	"github.com/go-asset/generics/list"
	"github.com/go-asset/generics/maybe"
)

func add[L any, R constraints.Ordered](c maybe.Maybe[R], value either.Either[L, R]) maybe.Maybe[R] {
	if !value.IsRight() {
		log.Printf("Skip: %v\n", value.Left())

		return c
	}

	if c.IsNothing() {
		return maybe.Just(value.Right())
	}

	return maybe.Just(c.Value() + value.Right())
}

func main() {
	result := list.FoldlIter(
		maybe.Nothing[int](),
		readFromUser(),
		add[string, int],
	)

	fmt.Println(result)
}
