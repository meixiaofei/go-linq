package demo

import (
	"fmt"
	"go-linq/linq"
	"reflect"
	"testing"
)

func Test_Linq(t *testing.T) {
	var arrs = []int{1, 2, 36, 78, 9}
	var b = linq.AnyComparable(arrs, 1, 3)
	println(b)
	var b2 = linq.AnyComparable(arrs, 4, 3)
	println(b2)

	var b3 = linq.AnyComparableChild(arrs, 1, 3)
	println(b3)

	var c = Dic[KeyPair[int, int]]{
		El: KeyPair[int, int]{
			Key:   1,
			Value: 2,
		},
	}
	println(c.El.Key)

	var d = KeyPair[int, int]{
		Key:   1,
		Value: 2,
	}
	println(d.Value)

	var e = Dic[int]{
		El: 1,
	}
	println(e.El)
}

func Test_Reflect(t *testing.T) {
	hl := hello[string]
	fv := reflect.ValueOf(hl)
	fmt.Println("fv is reflect.Func ?", fv.Kind() == reflect.Func)
	fv.Call(nil)
}

func Test_Reflect2(t *testing.T) {
	myType := &MyType2{"asd"}
	mtV := reflect.ValueOf(&myType).Elem()
	params := make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf("OOOOOO"))
	mtV.MethodByName("Hello").Call(params)
}

func Test_Reflect3(t *testing.T) {
	myType := &MyType[string]{"asd"}
	mtV := reflect.ValueOf(&myType).Elem()
	params := make([]reflect.Value, 0)
	params = append(params, reflect.ValueOf("OOOOOO"))
	mtV.MethodByName("Hello").Call(params)
}

type MyType[T any] struct {
	name T
}

func (mt *MyType[T]) Hello(show string) {
	println(mt.name, show)
}

type MyType2 struct {
	name string
}

func (mt *MyType2) Hello(show string) {
	println(mt.name, show)
}

func hello[T any](v T) {
	println(v)
}

type KeyPair[K any, V any] struct {
	Key   K
	Value V
}

type Dic[E any] struct {
	El E
}
