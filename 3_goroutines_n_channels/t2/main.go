package main

import (
	"fmt"
	"time"
)

func foo() {
	for {
		fmt.Println("Hello")
		time.Sleep(100 * time.Millisecond)
	}
}

func main() {
	go foo()

	// запуск горутины при помощи анонимной функции:
	// * это нужно для того, чтобы запустить какой-то
	// один кусок кода, который больше не будет исполь-
	// зоваться нигде вообще
	go func() {
		for {
			fmt.Println("Anon")
			time.Sleep(100 * time.Millisecond)
		}
	}()

	time.Sleep(1 * time.Second)
}
