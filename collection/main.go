package main

import "fmt"

func main() {
	arrayTest()
	arrayInitTest()
	arrayMultidimentionalTest()

	sliceTest() // slice is just not method, it's flexible data structure like list
	sliceCopyTest()
	mapTest()
}

func arrayTest() {
	var a [3]int
	a[0] = 1
	a[1] = 2
	a[2] = 3
	fmt.Println(a[1])
}

func arrayInitTest() {
	// construct/initialze with values
	var b = [3]int{11, 22, 33}
	fmt.Println(b[1])
}

func arrayMultidimentionalTest() {
	// multi dimentionalal array test
	var a = [2][3]int{
		{1, 2, 3},
		{4, 5, 6},
	}

	fmt.Println(a[1][2])
}

func sliceTest() {
	var a []int // it's slice, not array if you no describe size

	a = []int{1, 2, 3}
	a[1] = 10
	fmt.Println("a", a, len(a), cap(a))

	s := make([]int, 5, 10)
	fmt.Println("s", s, len(s), cap(s))

	s = s[2:5] // sub-slice (--like python--)
	fmt.Println("s", s, len(s), cap(s))

	s = append(s, 99, 88) // extend size 2 with appending values 99, 88
	fmt.Println("s", s, len(s), cap(s))

}

func sliceCopyTest() {
	fmt.Println("-- sliceCopyTest")
	source := []int{0, 1, 2}
	target := make([]int, len(source), cap(source)*2)

	fmt.Println("source", source, len(source), cap(source))
	fmt.Println("target", target, len(target), cap(target))

	copy(target, source) // copy to target from source

	fmt.Println("source", source, len(source), cap(source))
	fmt.Println("target", target, len(target), cap(target))
}

func mapTest() {
	fmt.Println("--mapTest")
	tickers := map[string]string{
		"GOOG": "Google Inc",
		"MSFT": "Microsoft",
		"FB":   "Facebook",
		"AMZN": "Amazon",
	}
	fmt.Println(tickers)

	var m map[int]string
	m = make(map[int]string)
	m[901] = "Apple"
	m[134] = "Grape"
	m[777] = "Tomato"

	str := m[134]
	fmt.Println(str)

	noData := m[999]
	fmt.Println(noData)
	delete(m, 777)
	fmt.Println(m[777])

	_, exists := tickers["AZZZ"]

	if !exists {
		fmt.Println("No AZZZ ticker")
	}
}
