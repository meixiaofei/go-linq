package linq

import (
	"golang.org/x/exp/constraints"
	"golang.org/x/exp/slices"
)

type OrderEnumerable[E constraints.Ordered] struct {
	ComparableEnumerable[E]
}

func AsOrderEnumerable[E constraints.Ordered](arr []E) OrderEnumerable[E] {
	return OrderEnumerable[E]{AsComparableEnumerable(arr)}
}

func (e OrderEnumerable[E]) Max() E {
	return Max(e.ComparableEnumerable.values, func(e E) E { return e })
}

func (e OrderEnumerable[E]) Min() E {
	return Min(e.ComparableEnumerable.values, func(e E) E { return e })
}

func (e OrderEnumerable[E]) Sum() E {
	var sum E
	for _, v := range e.ComparableEnumerable.values {
		sum = sum + v
	}
	return sum
}

func (e OrderEnumerable[E]) Sort() OrderEnumerable[E] {
	slices.Sort(e.ComparableEnumerable.values)
	return e
}

func (e OrderEnumerable[E]) Where(predicate func(E) bool) OrderEnumerable[E] {
	return AsOrderEnumerable(Where(e.Enumerable.values, predicate))
}

func (e OrderEnumerable[E]) Select(mapper func(E) E) OrderEnumerable[E] {
	return AsOrderEnumerable(Select(e.Enumerable.values, mapper))
}

func (e OrderEnumerable[E]) Take(size int) OrderEnumerable[E] {
	return AsOrderEnumerable(Take(e.Enumerable.values, size))
}
