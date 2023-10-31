package linq

import (
	"fmt"
	"strings"

	"golang.org/x/exp/constraints"
)

func MapKeyToSlice[K comparable, V any](source map[K]V) []K {
	return MapToSlice(source, func(k K, v V) K { return k })
}

func MapValueToSlice[K comparable, V any](source map[K]V) []V {
	return MapToSlice(source, func(k K, v V) V { return v })
}

func MapToSlice[K comparable, V any, R any](source map[K]V, fun func(K, V) R) []R {
	var res = make([]R, 0, len(source))
	for key, v := range source {
		res = append(res, fun(key, v))
	}
	return res
}

func MapSelect[K comparable, V any, R any](source map[K]V, fun func(K, V) (K, R)) map[K]R {
	var res = make(map[K]R, len(source))
	for key, v := range source {
		newKey, newV := fun(key, v)
		res[newKey] = newV
	}
	return res
}

func MapWhere[K comparable, E any](values map[K]E, predicate func(K, E) bool) map[K]E {
	var res = make(map[K]E, len(values))
	for k, v := range values {
		if predicate(k, v) {
			res[k] = v
		}
	}
	return res
}

//MapWhereSelect  按照条件过滤筛选返回切片集合
func MapWhereSelect[K comparable, E any, R any](values map[K]E, predicate func(K, E) (R, bool)) []R {
	var res = make([]R, 0, len(values))
	for key, v := range values {
		if r, ok := predicate(key, v); ok {
			res = append(res, r)
		}
	}
	return res
}

func Distinct[E any, K comparable](values []E, fun func(E) K) []E {
	if values == nil && len(values) < 2 {
		return values
	}
	newSlice := make([]E, 0)
	distinct := map[K]struct{}{}
	for _, v := range values {
		key := fun(v)
		if _, ok := distinct[key]; ok {
			continue
		}
		distinct[key] = struct{}{}
		newSlice = append(newSlice, v)
	}
	return newSlice
}

func DistinctComparable[E comparable](values []E) []E {
	return Distinct(values, func(e E) E {
		return e
	})
}

func Select[S any, T any](source []S, fun func(S) T) []T {
	var res = make([]T, 0, len(source))
	for _, v := range source {
		res = append(res, fun(v))
	}
	return res
}

func Take[E any](values []E, size int) []E {
	if values == nil {
		return nil
	}
	newSlice := make([]E, 0, size)
	for i := 0; i < len(values) && i < size; i++ {
		newSlice = append(newSlice, values[i])
	}
	return newSlice
}

func First[E any](values []E, predicates ...func(E) bool) (e E) {
	for _, v := range values {
		for _, predicate := range predicates {
			if predicate(v) {
				return v
			}
		}
		if len(predicates) == 0 {
			return v
		}
	}
	return e
}

func Last[E any](values []E, predicates ...func(E) bool) (e E) {
	if len(values) == 0 {
		return e
	}
	for i := len(values) - 1; i >= 0; i-- {
		if len(predicates) == 0 {
			return values[i]
		}
		for _, predicate := range predicates {
			if predicate(values[i]) {
				return values[i]
			}
		}
	}

	return e
}

func Where[E any](values []E, predicate func(E) bool) []E {
	var res = make([]E, 0, len(values))
	for _, v := range values {
		if predicate(v) {
			res = append(res, v)
		}
	}
	return res
}

func All[E any](values []E, predicate func(E) bool) bool {
	for _, v := range values {
		if !predicate(v) {
			return false
		}
	}
	return true
}

func AllComparable[E comparable](values []E, e E) bool {
	return All(values, func(r E) bool { return e == r })
}

func Any[E any](values []E, predicate func(E) bool) bool {
	for _, v := range values {
		if predicate(v) {
			return true
		}
	}
	return false
}

//AnyComparable 查找元素是否存在，如果查找对象是个切片，就判断是否相交
func AnyComparable[E comparable](values []E, es ...E) bool {
	return Any(values, func(r E) bool {
		return Any(es, func(e E) bool {
			return e == r
		})
	})
}

//AnyComparableChild 是否存在子集
func AnyComparableChild[E comparable](values []E, es ...E) bool {
	return All(es, func(e E) bool {
		return Any(values, func(v E) bool {
			return e == v
		})
	})
}

func Join[E any](values []E, fun func(E) string, sep string) string {
	if len(values) == 0 {
		return ""
	}
	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = fun(v)
	}
	return strings.Join(strs, sep)
}

func JoinOrdered[E constraints.Ordered](values []E, sep string) string {
	if len(values) == 0 {
		return ""
	}
	strs := make([]string, len(values))
	for i, v := range values {
		strs[i] = fmt.Sprintf("%d", v)
	}
	return strings.Join(strs, sep)
}

func Sum[E any, R constraints.Ordered](values []E, fun func(E) R) R {
	var sum R
	for _, v := range values {
		sum = sum + fun(v)
	}
	return sum
}

func SumOrdered[E constraints.Ordered](arr ...E) E {
	return Sum(arr, func(e E) E { return e })
}

func MaxOrdered[E constraints.Ordered](arr ...E) E {
	return Max(arr, func(e E) E { return e })
}

func MinOrdered[E constraints.Ordered](arr ...E) E {
	return Min(arr, func(e E) E { return e })
}

func Max[E any, R constraints.Ordered](arr []E, fun func(E) R) E {
	var max E
	for i, v := range arr {
		if i == 0 || fun(v) > fun(max) {
			max = v
		}
	}
	return max
}

func Min[E any, R constraints.Ordered](arr []E, fun func(E) R) E {
	var min E
	for i, v := range arr {
		if i == 0 || fun(v) < fun(min) {
			min = v
		}
	}
	return min
}

func GroupBy[K comparable, V any](arr []V, keyFunc func(V) K) map[K][]V {
	dic := map[K][]V{}
	for _, v := range arr {
		key := keyFunc(v)
		dic[key] = append(dic[key], v)
	}
	return dic
}

func ToMap[K comparable, V any, R any](arr []V, keyFunc func(V) K, vFunc func(K, []V) R) map[K]R {
	dic := map[K][]V{}
	for _, v := range arr {
		key := keyFunc(v)
		dic[key] = append(dic[key], v)
	}
	res := map[K]R{}
	for key, v := range dic {
		res[key] = vFunc(key, v)
	}
	return res
}

func ComparableAll[E comparable](values []E, e E) bool {
	return All(values, func(r E) bool { return r == e })
}

func ComparableAny[E comparable](values []E, e E) bool {
	return Any(values, func(r E) bool { return r == e })
}

func ToComparableMap[E comparable, V any](arr []E, vFunc func(E, []E) V) map[E]V {
	return ToMap(arr, func(v E) E { return v }, vFunc)
}

func ToBoolMap[E comparable](arr []E) map[E]bool {
	return ToComparableMap(arr, func(v E, es []E) bool { return true })
}

func AddMapKeys[E comparable](maps map[E]bool, keys ...E) {
	for _, key := range keys {
		maps[key] = true
	}
}

//ExchangeMap 交换key value
func ExchangeMap[K comparable, V comparable](maps map[K]V) map[V]K {
	dic := make(map[V]K, len(maps))
	for key, value := range maps {
		dic[value] = key
	}
	return dic
}

func If[T any](flag bool, f1, f2 T) T {
	if flag {
		return f1
	}
	return f2
}

func IfFunc[T any](flag bool, f1, f2 func() T) T {
	if flag {
		return f1()
	}
	return f2()
}
