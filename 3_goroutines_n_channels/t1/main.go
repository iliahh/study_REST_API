package main

import (
	"fmt"
	"time"
)

// Функция, описывающая поход рабочего в вахту
// Принимает:
// 1. Номер похода в шахту
// Возвращает:
// 2. Добытый уголь

func mine(transferPoint chan int, n int) {
	fmt.Println("Поход в шахту номер", n, "начался...")
	time.Sleep(1 * time.Second)
	fmt.Println("Поход в шахту номер", n, "закончился")

	transferPoint <- 10
	fmt.Println("Поход номер", n, "уголь передал!")
}

func main() {
	coal := 0

	// небуферизированная реализация
	// transferPoint := make(chan int)

	// буферизированная реализация
	transferPoint := make(chan int, 3)

	initTime := time.Now()

	// линейный поток выполнения, идущий друг за другом,
	// что тратит много времени
	/*
		coal += mine(1)
		coal += mine(2)
		coal += mine(3)
	*/

	// в данном случае задачи выполняются параллельно,
	// шахтёры добывают уголь не друг за другом в очереди,
	// а параллельно
	go mine(transferPoint, 1)
	go mine(transferPoint, 2)
	go mine(transferPoint, 3)

	coal += <-transferPoint
	time.Sleep(1 * time.Second)

	coal += <-transferPoint
	time.Sleep(1 * time.Second)

	coal += <-transferPoint
	time.Sleep(1 * time.Second)

	fmt.Println("Добыли", coal, "угля!")
	fmt.Println("Прошло времени:", time.Since(initTime))

	// канал не блокируется только когда,
	// в канал либо записывают что-то,
	// либо читают что-то, то есть это место синхронизации
	// работы двух сторон
	// - прием и передача должны быть активны

	// если мейн горутина завершается, то все остальные
	// завершаются также, поэтому, если отправить еще раз
	// после отработки 3 шахтеров забрать уголь, программа
	// выдаст deadlock.

	// утечки горутин не так страшны, как дедлоки,
	// но их тоже нужно избегать
}
