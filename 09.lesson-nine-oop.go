package main

import (
	"fmt"
)

func lessonNineOop() {
	// creates an object on local execution (stack)
	foo := myStruct{}
	foo.myField = "bar"
	fmt.Println(foo)
	fmt.Println(foo.myField)

	foo2 := myStruct{"bar2"} // order matters
	fmt.Println(foo2)
	fmt.Println(foo2.myField)

	// foo is a memory address now
	// better to use especially in case of large objects
	foo3 := new(myStruct) // reference type (heap)
	foo3.myField = "bar3"
	// no need to dereference
	fmt.Println(foo3)
	fmt.Println(foo3.myField)

	// shorter way of creating reference types
	foo4 := &myStruct{"bar4"}
	fmt.Println(foo4)
	fmt.Println(foo.myField)
	// go doesn't clean reference type if the memory reference is used outside local function
	// instead it keeps the record memory safe as long as it's needed

	// no constructors in go (constructor functions instead)
	foo5 := newMyStruct2()
	foo5.myMap["bar"] = "baz"
	fmt.Println(foo5)
}

// no classes, instead structs
type myStruct struct {
	myField string
}

type myStruct2 struct {
	// we need to init this before using
	myMap map[string]string
}

// constructor function, name is a convention
func newMyStruct2() *myStruct2 {
	result := myStruct2{}
	result.myMap = map[string]string{}

	//& <- return reference to result
	return &result
}
