package linq

import "golang.org/x/exp/slices"

type SelectEnumerable[S any, T any] struct {
	Enumerable[S]
}

func AsSelectEnumerable[S any, T any](arr []S) SelectEnumerable[S, T] {
	return SelectEnumerable[S, T]{AsEnumerable(arr)}
}

func (e SelectEnumerable[S, T]) Where(predicate func(S) bool) SelectEnumerable[S, T] {
	return AsSelectEnumerable[S, T](Where(e.values, predicate))
}

func (e SelectEnumerable[S, T]) Take(size int) SelectEnumerable[S, T] {
	return AsSelectEnumerable[S, T](Take(e.values, size))
}

func (e SelectEnumerable[S, T]) Sort(less func(a, b S) bool) SelectEnumerable[S, T] {
	slices.SortFunc(e.values, less)
	return e
}

func (e SelectEnumerable[S, T]) Select(mapper func(S) T) SelectEnumerable[T, S] {
	var res = make([]T, 0, len(e.values))
	for _, v := range e.values {
		res = append(res, mapper(v))
	}
	return AsSelectEnumerable[T, S](res)
}
