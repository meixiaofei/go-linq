package linq

type GroupEnumerable[K comparable, V any, R any] struct {
	Groups []Group[K, V]
}

type Group[K comparable, V any] struct {
	Enumerable[V]
	Key K
}

func AsGroupEnumerable[K comparable, V any, R any](arr []V, fun func(V) K) GroupEnumerable[K, V, R] {
	dic := GroupBy(arr, fun)
	var res = make([]Group[K, V], 0)
	for key, v := range dic {
		res = append(res, Group[K, V]{AsEnumerable(v), key})
	}
	return GroupEnumerable[K, V, R]{
		Groups: res,
	}
}

func (e GroupEnumerable[K, V, R]) Where(predicate func(K, Enumerable[V]) bool) GroupEnumerable[K, V, R] {
	var res = make([]Group[K, V], 0)
	for _, v := range e.Groups {
		if predicate(v.Key, v.Enumerable) {
			res = append(res, v)
		}
	}
	return GroupEnumerable[K, V, R]{
		Groups: res,
	}
}

func (e GroupEnumerable[K, V, R]) Take(size int) GroupEnumerable[K, V, R] {
	if e.Groups == nil {
		return e
	}
	newSlice := make([]Group[K, V], 0, size)
	for i := 0; i < len(e.Groups) && i < size; i++ {
		newSlice = append(newSlice, e.Groups[i])
	}
	return GroupEnumerable[K, V, R]{
		Groups: newSlice,
	}
}

func (e GroupEnumerable[K, V, R]) Select(mapper func(Group[K, V]) R) Enumerable[R] {
	return AsEnumerable[R](Select(e.Groups, mapper))
}

func (e GroupEnumerable[K, V, R]) ToMap(mapFunc func(Group[K, V]) R) map[K]R {
	res := map[K]R{}
	for _, v := range e.Groups {
		res[v.Key] = mapFunc(v)
	}
	return res
}
