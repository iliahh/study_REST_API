package main

import (
	"fmt"
	"time"
)

// Условие задачи по анонимным функциям в горутинах:
// - запустите 3 горутины через различные анонимные функции
// - каждая анонимная функция должна 3 раза выводить на экран любой текст
// - посмотрите на консольнный вывод, сравните этот вывод с вашими ожиданиями

func main() {
	go func() {
		for i := 1; i <= 4; i++ {
			fmt.Printf("Анонимная функция №1 выводится %d раз/а\n", i)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Анонимная функция №2 выводится %d раз/а\n", i)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	go func() {
		for i := 1; i <= 3; i++ {
			fmt.Printf("Анонимная функция №3 выводится %d раз/а\n", i)
			time.Sleep(250 * time.Millisecond)
		}
	}()

	time.Sleep(3 * time.Second)
}
