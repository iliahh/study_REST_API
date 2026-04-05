package main

import (
	"fmt"
	"time"
)

func main() {
	intCh := make(chan int)
	strCh := make(chan string)

	go func() {
		time.Sleep(200 * time.Millisecond)
		intCh <- 1

		/*
			i := 1

			for {
				intCh <- i
				i++

				time.Sleep(10 * time.Second)
			}
		*/
	}()

	go func() {
		time.Sleep(100 * time.Millisecond)
		strCh <- "hi"

		/*
			i := 1
			for {
				strCh <- "hi" + strconv.Itoa(i)
				i++

				time.Sleep(1 * time.Second)
			}
		*/
	}()

	time.Sleep(50 * time.Millisecond)

	select {
	case number := <-intCh:
		fmt.Println("intCh:", number)
	case str := <-strCh:
		fmt.Println("strCh:", str)
	default:
		fmt.Println("Никакой канал не готов")
	}

	/*
		for {
			select {
			case number := <-intCh:
				fmt.Println("intCh:", number)
			case str := <-strCh:
				fmt.Println("strCh:", str)
			default:

			}

				// number := <-intCh
				// fmt.Println("intCh:", number)

				// str := <-strCh
				// fmt.Println("strCh:", str)
		}
	*/
}
