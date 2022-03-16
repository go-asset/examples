package main

import "github.com/go-asset/generics/either"

func readFromUser() chan either.Either[string, int] {
	ch := make(chan either.Either[string, int])

	myList := []either.Either[string, int]{
		either.Right[string](12),
		either.Right[string](2),
		either.Left[string, int]("asd"),
		either.Right[string](8),
	}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	return ch
}
