package main

import (
	"fmt"
)

func lessonElevenComposition() {
	// in go there is no inheritance, only composition
	// instead we should use constructor function for enhancedMessagePrinter
	myMp2 := enhancedMessagePrinter{messagePrinter2{"foo"}}
	myMp2.printMessage()
}

type messagePrinter2 struct {
	message string
}

func (mp2 *messagePrinter2) printMessage() {
	fmt.Println(mp2.message)
}

type enhancedMessagePrinter struct {
	messagePrinter2
}