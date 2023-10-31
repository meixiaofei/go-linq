package demo

import (
	"fmt"
	"go-linq/linq"
	"strconv"
	"testing"
)

func init() {
	s1 := linq.AsEnumerable([]string{"d", "a", "b", "c", "a"}).
		Where(func(s string) bool { return s != "b" }).
		Select(func(s string) string { return "class_" + s }).
		Sort(func(s1, s2 string) bool { return s1 < s2 }).
		ToSlice()
	fmt.Println(s1)

	s2 := linq.AsComparableEnumerable([]string{"d", "a", "b", "c", "a"}).
		Where(func(s string) bool { return s != "b" }).
		Select(func(s string) string { return "class_" + s }).
		Distinct().
		Take(3).
		ToSlice()
	fmt.Println(s2)

	s3 := linq.AsOrderEnumerable([]string{"d", "a", "b", "c", "a"}).
		Where(func(s string) bool { return s != "b" }).
		Select(func(s string) string { return "class_" + s }).
		Sort().
		Distinct().
		Take(3).
		ToSlice()
	fmt.Println(s3)

	s5 := linq.AsSelectEnumerable[int, string]([]int{1, 2, 3, 4, 5}).
		Where(func(v int) bool { return v > 3 }).
		Select(func(v int) string { return "mapping_" + strconv.Itoa(v) }).
		Take(2).
		ToSlice()

	fmt.Println(s5)

}

func TestLinq(t *testing.T) {
	var a = struct {
		Id   int
		Name string
	}{}
	Get(a) //编译不通过

	fmt.Println("开始")
}

func Get[T any](a T) {
}
