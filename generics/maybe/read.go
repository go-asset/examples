package main

import "github.com/go-asset/generics/maybe"

func readFromUser() chan maybe.Maybe[int] {
	ch := make(chan maybe.Maybe[int])

	myList := []maybe.Maybe[int]{
		maybe.Just(12),
		maybe.Just(2),
		maybe.Nothing[int](),
		maybe.Just(8),
	}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	return ch
}

func readFromUserStr() chan maybe.Maybe[string] {
	ch := make(chan maybe.Maybe[string])

	myList := []maybe.Maybe[string]{
		maybe.Just("asd"),
		maybe.Just("123"),
		maybe.Nothing[string](),
		maybe.Just(",./"),
	}

	go func() {
		for _, v := range myList {
			ch <- v
		}

		close(ch)
	}()

	return ch
}
