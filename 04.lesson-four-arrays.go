package main

import "fmt"

func lessonFourArrays() {
	myArray := [3]int{}
	myArray[0] = 42
	myArray[1] = 27
	myArray[2] = 99

	myArray2 := [...]int{1, 2, 3}

	fmt.Println(myArray)
	fmt.Println(myArray2)
	fmt.Println(len(myArray))

	mySlice := myArray[0:2]
	mySlice = append(mySlice, 100)
	fmt.Println(mySlice)

	//inefficient when adding a lot of elements
	mySlice2 := []float32{14., 15., 19.}
	fmt.Println(mySlice2)
	fmt.Println(len(mySlice2))

	// better when not resizing array all the time
	mySlice3 := make([]float32, 100)
	mySlice3[0] = 12.
	mySlice3[1] = 16.
	mySlice3[2] = 18.

	fmt.Println(mySlice3)
	fmt.Println(len(mySlice3))

	//map of key=int value=string
	myMap := make(map[int]string)
	fmt.Println(myMap)
	myMap[43] = "Foo"
	myMap[12] = "bar"

	fmt.Println(myMap)
	fmt.Println(myMap[99])
}
