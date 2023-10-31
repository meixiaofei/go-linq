package demo

import (
	"fmt"
	"github.com/meixiaofei/go-linq/linq"
	"strings"
	"testing"
)

type M struct {
	Score int
	Name  string
	Age   int
}

func TestGroup(t *testing.T) {
	var arr = []M{
		{30, "张三", 18},
		{79, "李四", 18},
		{46, "王五", 20},
		{88, "赵六", 20},
		{12, "熊七", 20},
	}

	//// 根据年龄分组，获取每组第一个人
	//var s1 = linq.AsGroupEnumerable[int, M, string](arr, func(m M) int { return m.Age }).
	//	ToMap(func(g linq.Group[int, M]) string { return g.Enumerable.First().Name })
	//fmt.Println(s1)
	//
	////根据年龄分组，获取每组Score最大的人
	var s2 = linq.AsGroupEnumerable[int, M, string](arr, func(m M) int { return m.Age }).
		ToMap(func(g linq.Group[int, M]) string { return g.Enumerable.MaxByInt(func(m M) int { return m.Score }).Name })
	fmt.Println(s2)

	//根据年龄分组，获取每组人数
	var s3 = linq.ToMap(arr,
		func(e M) int { return e.Age },
		func(key int, es []M) int { return len(es) })
	fmt.Println(s3)

	var words = "how do you do"
	//var countWords = words.Split(' ').GroupBy(x => x).ToDictionary(x => x.Key, x => x.Count());
	var s4 = linq.ToMap(strings.Split(words, " "),
		func(e string) string { return e },
		func(key string, es []string) int { return len(es) })
	fmt.Println(s4)

	//var s5 = linq.ToComparableMap(strings.Split(words, ","), func(key string, es []string) int { return len(es) })
	//fmt.Println(s5)
}
