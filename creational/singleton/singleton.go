package main

import (
	"fmt"
	"sync"
)

type printer struct{}

var instance *printer
var once sync.Once

func GetPrinter() *printer {
	once.Do(func() {
		instance = &printer{}
	})
	return instance
}

func (p *printer) PrintSentence(sentence string) {
	fmt.Println(sentence)
}
