package main

import "fmt"

// Закрытие каналов. Аксиомы каналов
func main() {
	// var ch chan int - nil канал

	// Создаём новый открытый канал
	// ch := make(chan int)

	// Закрываем канал
	// close(ch)

	// Аксиома 1:
	// Закрываю закрытый канал
	// close(ch) // panic: close of closed channel
	// - если мы попытаемся закрыть уже закрытый канал, то мы получим панику (согласно схеме)
	// (два раза мы не можем сигнализировать, что канал закрыт)

	// Аксиома 2:
	// Читаю значение из закрытого канала
	// v1 := <-ch
	// fmt.Println("v:", v1) // v: 0
	// - если мы читаем из закрытого канала,
	// то мы получаем значения по умолчанию для типа данных канала (в данном случае 0 для int),
	// также можем читать сколько угодно раз (v1, v2, v3 ...)

	// Аксиома 3:
	// Записываю значение в закрытый канал
	// ch <- 10 // panic: send on closed channel

	// Следующий пример:
	/*
		transferPoint := make(chan int)

		// шахтёр (добывает и отдаёт)
			go func() {
				// 3 4 5 6
				iterations := 3 + rand.Intn(4)
				fmt.Println("iterations:", iterations)

				for i := 1; i <= iterations; i++ {
					transferPoint <- 10
					time.Sleep(300 * time.Millisecond)
				}

				close(transferPoint) // Закрытие канала - БАЗА, необходимая для
				// корректной организации параллельной работы приложения
			}()
	*/

	// прораб (получает и считает)
	/*
		coal := 0
	*/

	// for {
	// 	v, ok := <-transferPoint
	// 	if !ok {
	// 		fmt.Println("Шахтёр обработал всю шахту!")
	// 		break
	// 	}

	// 	coal += v

	// 	fmt.Println("coal:", coal)
	// }

	// Ещё один способ - использовать цикл for range,
	// который автоматически обрабатывает закрытие канала
	// (ниже представлена короткая запись того, что выше под комментом)
	/*
		for v := range transferPoint {
			coal += v
			fmt.Println("coal:", coal)
		}

		fmt.Println("Суммарно добытый уголь:", coal)
	*/

	// Пример с nil-каналами:
	// (короче, мы не можем с nil-канала читать, писать и не можем его закрывать) ->
	// значит каналы всегда нужно инициализировать через = make(chan ...) !!!

	// var ch chan string
	var ch chan string = make(chan string) // вывод: v: Hello

	// Сделаем горутину через анон функцию
	go func() {
		ch <- "Hello" // одна горутина заблокировалась при попытке что-то записать в nil-канал
	}()

	v := <-ch // другая заблокировалась при попытке что-то в nil-канале прочитать

	fmt.Println("v:", v) // fatal error: all goroutines are asleep - deadlock!

	// close(ch) // panic: close of nil channel
}
