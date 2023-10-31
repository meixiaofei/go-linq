package linq

import "golang.org/x/exp/slices"

type ComparableEnumerable[E comparable] struct {
	Enumerable[E]
}

func AsComparableEnumerable[E comparable](arr []E) ComparableEnumerable[E] {
	return ComparableEnumerable[E]{AsEnumerable(arr)}
}

func (e ComparableEnumerable[E]) Distinct() ComparableEnumerable[E] {
	if e.values == nil && len(e.values) < 2 {
		return e
	}
	newSlice := make([]E, 0)
	distinct := map[E]struct{}{}
	for _, v := range e.values {
		if _, ok := distinct[v]; ok {
			continue
		}
		distinct[v] = struct{}{}
		newSlice = append(newSlice, v)
	}
	e.values = newSlice
	return e
}

func (e ComparableEnumerable[E]) Where(predicate func(E) bool) ComparableEnumerable[E] {
	return AsComparableEnumerable(e.Enumerable.Where(predicate).values)
}

func (e ComparableEnumerable[E]) Select(mapper func(E) E) ComparableEnumerable[E] {
	return AsComparableEnumerable(e.Enumerable.Select(mapper).values)
}

func (e ComparableEnumerable[E]) Take(size int) ComparableEnumerable[E] {
	return AsComparableEnumerable(e.Enumerable.Take(size).values)
}

func (e ComparableEnumerable[E]) SortFunc(less func(a, b E) int) ComparableEnumerable[E] {
	slices.SortFunc(e.values, less)
	return e
}
