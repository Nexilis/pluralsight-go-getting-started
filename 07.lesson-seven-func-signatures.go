package main

func lessonSevenFunctionSignatures() {
	message := "Hello"
	sayHelloWithValueCopy(message)
	println(message)

	message2 := "Gello"
	sayHelloUsingReference(&message2)
	println(message2)

	sayHelloVariadic("Hello", "Go", "from", "Pluralsight")
}

func sayHelloWithValueCopy(msg string) {
	println(msg)

	msg = "Hello Go"
}

func sayHelloUsingReference(msg *string) {
	println(*msg)

	*msg = "Gello Ho"
}

func sayHelloVariadic(messages ...string) {
	for _, message := range messages {
		println(message)
	}
}