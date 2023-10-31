package linq

/*
备注改文件是因为Go泛型不支持方法，临时方案
*/

func (e Enumerable[E]) MaxByString(fun func(E) string) E {
	return Max(e.values, fun)
}

func (e Enumerable[E]) MaxByInt(fun func(E) int) E {
	return Max(e.values, fun)
}

func (e Enumerable[E]) MaxByInt64(fun func(E) int64) E {
	return Max(e.values, fun)
}

func (e Enumerable[E]) MaxByFloat64(fun func(E) float64) E {
	return Max(e.values, fun)
}

func (e Enumerable[E]) MaxByFloat32(fun func(E) float32) E {
	return Max(e.values, fun)
}

func (e Enumerable[E]) MinByString(fun func(E) string) E {
	return Min(e.values, fun)
}

func (e Enumerable[E]) MinByInt(fun func(E) int) E {
	return Min(e.values, fun)
}

func (e Enumerable[E]) MinByInt64(fun func(E) int64) E {
	return Min(e.values, fun)
}

func (e Enumerable[E]) MinByFloat64(fun func(E) float64) E {
	return Min(e.values, fun)
}

func (e Enumerable[E]) MinByFloat32(fun func(E) float32) E {
	return Min(e.values, fun)
}

func (e Enumerable[E]) SelectString(fun func(E) string) Enumerable[string] {
	return AsEnumerable(Select(e.values, fun))
}

func (e Enumerable[E]) SelectInt(fun func(E) int) Enumerable[int] {
	return AsEnumerable(Select(e.values, fun))
}

func (e Enumerable[E]) SelectInt64(fun func(E) int64) Enumerable[int64] {
	return AsEnumerable(Select(e.values, fun))
}
