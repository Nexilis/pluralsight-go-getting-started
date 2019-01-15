package main

const (
	first = iota
	second
	third
)

const (
	fourth = iota
)

const (
	one = 1 << (10 * iota)
	two
	three
)

func lessonThreeConsts() {
	println(first)
	println(second)
	println(third)
	println(fourth)

	println(one)
	println(two)
	println(three)
}
