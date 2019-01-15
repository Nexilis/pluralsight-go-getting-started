package main

import "fmt"

const (
	message = "Hello Go!"
)

var (
	message2 string = "Hello ple ple!"
	message3 string
)

func lessonOneVariables() {

	fmt.Println(message)
	message2 = "mutate me"
	fmt.Println(message2)
	fmt.Println(message3)
}

func init() {
	message3 = "Value from init"
}
