package main

func main() {

	basicTest()
	highorderTest()
	closureTest()
}

func basicTest() {
	sum := func(n ...int) int {
		s := 0
		for _, i := range n {
			s += i
		}
		return s
	}

	result := sum(1, 2, 3, 4, 5)
	println(result)
}

func highorderTest() {
	add := func(i int, j int) int {
		return i + j
	}

	r1 := calc(add, 10, 20)
	println(r1)

	r2 := calc(func(x int, y int) int { return x - y }, 10, 20)
	println(r2)

}

func calc(f func(int, int) int, a int, b int) int {
	result := f(a, b)
	return result
}

func closureTest() {

	next := nextValue()
	println(next())
	println(next())
	println(next())

	anotherNext := nextValue()
	println(anotherNext()) // 1
	println(anotherNext()) // 2
}

func nextValue() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
