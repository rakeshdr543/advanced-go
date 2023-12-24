package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func add[T int | float64](a T, b T) T {
	return a + b
}

type userId int

type intType interface{}

func add2[T constraints.Ordered](a, b T) T {
	return a + b
}

func MapValues[T comparable](values []T, multiplier func(T) T) []T {
	newValues := []T{}
	for _, v := range values {
		newValue := multiplier(v)
		newValues = append(newValues, newValue)
	}
	return newValues
}

type CustomData interface {
	constraints.Ordered | []byte | []rune
}

type User[T CustomData] struct {
	email string
	data  T
}

type Mapper[K comparable, V int | string] map[K]V

func main() {
	// res := add(3, 4)
	// fmt.Println("add integer", res)
	// res2 := add(1.3, 1.4)
	// fmt.Println("add float64", res2)
	// res3 := add2(3, 4)
	// fmt.Println("add integer", res3)
	// res4 := add2(1.3, 1.4)
	// fmt.Println("add float64", res4)

	// res := MapValues([]int{1, 2, 3}, func(a int) int {
	// 	return a * 10
	// })
	// fmt.Println(res)

	// newUser := User[string]{
	// 	email: "user@example.com",
	// 	data:  "admin",
	// }
	// newUser := User[int]{
	// 	email: "user@example.com",
	// 	data:  123,
	// }
	// fmt.Println(newUser)

	newMap := make(Mapper[string, int])
	newMap["3"] = 3
	fmt.Println(newMap)
}
