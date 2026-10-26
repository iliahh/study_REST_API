package main

/*
- создайте 3 канала (chan int, chan string, chan float64)
- создайте 3 горутины
- каждая горутина принимает один из созданных каналов,
после чего в бесконечном цикле начинает что-то в этот самый канал писать.
Интервал между итерациями 300ms, 1s, 5s для каждой горутины соответственно.
- в main горутине необходимо вычитывать приходящие в 3 канала значения
и выводить их на экран, при этом мы не должны блокироваться на ожидании
какого-то конкретного канала (а вдруг в другие каналы уже пришло значение,
а мы всё ждём один конкретный канал?)
- как обычно, оценить консольный вывод, понять что происходит в программе,
и сравнить это с ожиданиями
*/

import (
	"fmt"
	"time"
)

func main() {
	firstChan := make(chan int)
	secondChan := make(chan string)
	thirdChan := make(chan float64)

	go func() {
		for {
			firstChan <- 67

			time.Sleep(300 * time.Millisecond)
		}
	}()

	go func() {
		for {
			secondChan <- "жомдбе"

			time.Sleep(1 * time.Second)
		}
	}()

	go func() {
		for {
			thirdChan <- 6.767

			time.Sleep(5 * time.Second)
		}
	}()

	/*
		fc := <-firstChan
		fmt.Println(fc)

		sc := <-secondChan
		fmt.Println(sc)

		tc := <-thirdChan
		fmt.Println(tc)
	*/

	for {
		select {
		case fc := <-firstChan:
			fmt.Println(fc)
		case sc := <-secondChan:
			fmt.Println(sc)
		case tc := <-thirdChan:
			fmt.Println(tc)
		}
	}
}
