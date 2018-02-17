package main

import "fmt"

const (
	Apple = iota
	Grape
	Orange
)

func main() {

	basicTest()
	ifTest()
	switchTest()
	switchTest2()
	forTest()
	funcTest()
}

func basicTest() {
	var name string = "park minu"
	var name_p = &name
	var age int = 10

	*name_p = "minostar"

	var say string = fmt.Sprintf("hello, %s! are you %v?", name, age)
	println(say)
}

func ifTest() {
	k := int32(1)
	if k == 1 {
		println("One")
	} else if k == 2 {
		println("Two")
	} else {
		println("Other")
	}
}

func switchTest() {

	var category = 1
	var name string

	switch category {
	case 1:
		name = "paper book"
	case 2:
		name = "ebook"
	case 3, 4: // multiple match!
		name = "blog"
	default:
		name = "others"
	}
	println(name)

	// switch statement-expression!
	switch x := category << 2; x - 1 {
	case 1:
		name = "hi"
	case 2:
		name = "wow"
	}
	println(name)
}

func switchTest2() {

	score := 85

	// case expression!
	switch {
	case score >= 90:
		println("A")
	case score >= 80:
		println("B")
		fallthrough // don't break automatically!
	case score >= 70:
		println("C")
	case score >= 60:
		println("D")
	default:
		println("No Hope")
	}

	// switch what.(type) {
	// case int:
	// 	println("int")
	// case bool:
	// 	println("bool")
	// case string:
	// 	println("string")
	// default:
	// 	println("unknown")
	// }
}

func forTest() {
	names := []string{"hong", "lee", "kim"}

	for index, name := range names {
		println(index, name)
	}
}

func funcTest() {
	sum, multiplies := sum_and_multiplies(1, 3, 5, 7, 9)
	println(sum, multiplies)
}

// 1. return tuple(named return value)
// 2. use variadic function (variadic parameters)
func sum_and_multiplies(nums ...int) (int, int) {
	s := 0
	m := 1

	for _, n := range nums {
		s += n
		m *= n
	}

	return s, m
}
