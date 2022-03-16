package filter

import (
	"fmt"
	"strings"

	"github.com/go-asset/generics/list"
)

type Filter interface {
	String() string
}

type NameFilter struct {
	value string
}

func (f NameFilter) String() string {
	return fmt.Sprintf("name=%s", f.value)
}

func Name(value string) NameFilter {
	return NameFilter{value: value}
}

type AgeFilter struct {
	value int
}

func (f AgeFilter) String() string {
	return fmt.Sprintf("name=%d", f.value)
}

func Age(value int) AgeFilter {
	return AgeFilter{value: value}
}

type OrFilter struct {
	values []Filter
}

func (f OrFilter) String() string {
	values := list.Foldl([]string{}, f.values, func(c []string, value Filter) []string {
		return append(c, value.String())
	})

	return fmt.Sprintf("(%s)", strings.Join(values, " || "))
}

func Or(values ...Filter) OrFilter {
	return OrFilter{values: values}
}

type AndFilter struct {
	values []Filter
}

func (f AndFilter) String() string {
	values := list.Foldl([]string{}, f.values, func(c []string, value Filter) []string {
		return append(c, value.String())
	})

	return fmt.Sprintf("(%s)", strings.Join(values, " && "))
}

func And(values ...Filter) AndFilter {
	return AndFilter{values: values}
}
