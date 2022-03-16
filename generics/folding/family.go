package main

type Person struct {
	Name string
	Age  int
}

type Family struct {
	Adults   []Person
	Children []Person
}

func newFamily() Family {
	return Family{Adults: []Person{}, Children: []Person{}}
}
