package main

import "fmt"

/*
Домашнее задание по теме "Закрытие каналов. Аксиомы каналов".
- опробовать все рассмотренные аксиомы каналов собственными руками
*/

func main() {
	// Аксиома №1: закрытие закрытого канала
	/*
		// канал без буфера синхронизирует main() горутину и горутину отправки:
		// отправка завершится только тогда, когда значение будет прочитано.
		ch := make(chan int) // канал без буфера

		// запись идет через горутину, потому что иначе, если отправить и читать в одной
		// горутине по очереди, главная горутина остановится на отправке, ожидая чтения,
		// но сама горутина не дошла до чтения, в итоге получаем
		// all goroutines are asleep - deadlock!
		go func() {
			ch <- 10 // отправка десяточки, в канале без буфера отправка ждёт получателя
		}()

		value := <-ch // чтение десяточки; и когда выполняется эта строка,
		// то отправка и чтение встречаются, тогда 10 передаётся, и обе операции завершаются
		fmt.Println(value)

		close(ch)
		close(ch)
	*/

	/*
		Вывод:
		10
		panic: close of closed channel

		goroutine 1 [running]:
		main.main()
		        H:/GO/Study_REST_API/4_select/homework3/main.go:28 +0xd8
		exit status 2
	*/

	/*
		*or (2 var)*
		ch := make(chan int, 1) // канал с буфером на 1 значение
		ch <- 10 // значение попадает в буфер, отправка завершается
		close(ch) // закрываем канал, значение остаётся в буфере

		value := <-ch // читаем из буфера
		fmt.Println(value) // десяточка
	*/

	// Аксиома №2: чтение из закрытого канала
	/*
		// если читаем из закрытого канала, то мы получаем значение по умолчанию для
		// типа данных канала: int, float, complex - 0; string - "" (пустая строка); bool - false
		ch := make(chan int)

		go func() {
			ch <- 10
		}()

		close(ch)

		value := <-ch
		fmt.Println(value)
	*/

	/*
		Вывод:
		0
		panic: send on closed channel

		goroutine 20 [running]:
		main.main.func1()
		        H:/GO/Study_REST_API/4_select/homework3/main.go:59 +0x1e
		created by main.main in goroutine 1
		        H:/GO/Study_REST_API/4_select/homework3/main.go:58 +0x66
		exit status 2
	*/

	// Аксиома №3: запись значения в закрытый канал
	ch := make(chan int)
	close(ch)
	ch <- 10
	value := <-ch
	fmt.Println(value)

	/*
		Вывод:
		panic: send on closed channel

		goroutine 1 [running]:
		main.main()
		        H:/GO/Study_REST_API/4_select/homework3/main.go:86 +0x45
		exit status 2
	*/

	// закрытие канала - важный момент, это позволяет корректно организовать
	// параллельную работу приложения

	/*
		// автоматическое закрытие канала можно сделать через цикл for range:
		/*
		for v := range transferPoint {
			coal += v
			fmt.Println("coal:", coal)
		}

		fmt.Println("Суммарно добытый уголь:", coal)
	*/

	/*
		// мы не можем читать, писать и закрывать nil-каналы,
		// поэтому каналы всегда нужно инициализировать через make(chan ...)!

		var ch chan string = make(chan string) // вывод: v: Hello

		go func() {
			ch <- "Hello" // одна горутина заблокировалась при попытке что-то записать в nil-канал
		}()

		v := <-ch // другая заблокировалась при попытке что-то в nil-канале прочитать

		fmt.Println("v:", v) // fatal error: all goroutines are asleep - deadlock!

		// close(ch) // panic: close of nil channel
	*/
}
