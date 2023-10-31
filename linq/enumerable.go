package linq

import (
	"golang.org/x/exp/slices"
)

type values[E any] []E
type Enumerable[E any] struct {
	values[E]
}
type KeyPair[K any, V any] struct {
	Key   K
	Value V
}

func AsEnumerable[E any](arr []E) Enumerable[E] {
	return Enumerable[E]{arr}
}

func AsMapEnumerable[K comparable, V any](maps map[K]V) Enumerable[KeyPair[K, V]] {
	var arr = make([]KeyPair[K, V], 0, len(maps))
	for k, v := range maps {
		arr = append(arr, KeyPair[K, V]{Key: k, Value: v})
	}
	return Enumerable[KeyPair[K, V]]{arr}
}

func (e Enumerable[E]) All(predicate func(E) bool) bool {
	return All(e.values, predicate)
}

func (e Enumerable[E]) Any(predicate func(E) bool) bool {
	return Any(e.values, predicate)
}

func (e Enumerable[E]) Join(fun func(E) string, sep string) string {
	return Join(e.values, fun, sep)
}

func (e Enumerable[E]) ToSlice() []E {
	return e.values
}

func (e Enumerable[E]) Foreach(action func(int, E)) {
	if e.values == nil {
		return
	}
	for i := 0; i < len(e.values); i++ {
		action(i, e.values[i])
	}
}

func (e Enumerable[E]) Count() int {
	return len(e.values)
}

//First 获取第一个匹配的结果
func (e Enumerable[E]) First(predicates ...func(E) bool) (elem E) {
	if len(e.values) == 0 {
		return elem
	}
	if len(predicates) == 0 {
		return e.values[0]
	}

	for _, v := range e.values {
		for _, predicate := range predicates {
			if predicate(v) {
				return v
			}
		}
	}
	return elem
}

func (e Enumerable[E]) Where(predicate func(E) bool) Enumerable[E] {
	return AsEnumerable(Where(e.values, predicate))
}

func (e Enumerable[E]) Select(mapper func(E) E) Enumerable[E] {
	return AsEnumerable(Select(e.values, mapper))
}

func (e Enumerable[E]) Sort(less func(a, b E) bool) Enumerable[E] {
	slices.SortFunc(e.values, less)
	return e
}

func (e Enumerable[E]) Take(size int) Enumerable[E] {
	if e.values == nil {
		return e
	}
	newSlice := make([]E, 0, size)
	for i := 0; i < len(e.values) && i < size; i++ {
		newSlice = append(newSlice, e.values[i])
	}
	return AsEnumerable(newSlice)
}
