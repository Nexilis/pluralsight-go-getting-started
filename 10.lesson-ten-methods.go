package main

import (
	"fmt"
)

func lessonTenMethods() {
	myMp := messagePrinter{"foo"}
	myMp.printMessage()

}

type messagePrinter struct {
	message string
}

func (mp *messagePrinter) printMessage() {
	fmt.Println(mp.message)
}