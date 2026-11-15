package main

import (
	"concurrency/miner"
	"concurrency/postman"
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

func main() {
	// у нас есть гонка данных, которая не детектится даже флажком -race,
	// гонка происходит из-за доступа двух и более горутин (main и анонимная, где сам coal += v)
	// к одной переменной coal, ее нужно регулировать, в данном случае атомиками
	var coal atomic.Int64 // var coal int

	// пишем мы всё время (mails = append(mails, v)), а читаем единожды в самом конце,
	// поэтому конкурентный доступ будет регулироваться Mutex, а не RWMutex
	mtx := sync.Mutex{}
	var mails []string

	minerContext, minerCancel := context.WithCancel(context.Background())
	postmanContext, postmanCancel := context.WithCancel(context.Background())

	go func() {
		time.Sleep(3 * time.Second)
		fmt.Println("--->>  Рабочий день шахтёров окончен!")
		minerCancel()
	}()

	go func() {
		time.Sleep(6 * time.Second)
		fmt.Println("--->>  Рабочий день почтальонов окончен!")
		postmanCancel()
	}()

	coalTransferPoint := miner.MinerPool(minerContext, 2) // имя_пакета.Имя_функции_в_пакете
	mailTransferPoint := postman.PostmanPool(postmanContext, 2)

	initTime := time.Now()

	/*
		флажки с информацией о том, закрыт ли тот или иной канал для того,
		чтобы не читать дальше оттуда сообщения и для того, чтобы когда оба канала закрылись,
		я мог завершить наш цикл. получается, что цикл будет теперь не совсем бесконечным,
		и будет идти до тех пор, пока открыт канал для передачи угля или до тех пор,
		пока открыт канал для передачи писем, и когда я получу значение по умолчанию для ok,
		то есть канал для передачи угля закрыт, и я напишу, что isCoalClosed = true,
		а дальше я перейду на следующую итерацию.

		isCoalClosed := false
		isMailClosed := false
		for !isCoalClosed || !isMailClosed {
			select {
			case c, ok := <-coalTransferPoint:
				if !ok {
					isCoalClosed = true
					continue
				}
				coal += c
			case m, ok := <-mailTransferPoint:
				if !ok {
					isMailClosed = true
					continue
				}
				mails = append(mails, m)
			}
		}
	*/
	// кусок кода выше с 42 по -59 строку выглядит стрёмно, его можно представить иначе:
	wg := &sync.WaitGroup{}

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range coalTransferPoint {
			coal.Add(int64(v))
		}
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()

		for v := range mailTransferPoint {
			mtx.Lock()
			mails = append(mails, v)
			mtx.Unlock()
		}
	}()

	wg.Wait()

	fmt.Println("Суммарно добытый уголь:", coal.Load())

	mtx.Lock()
	fmt.Println("Суммарное количество полученных писем:", len(mails))
	mtx.Unlock()

	fmt.Println("Затраченное время:", time.Since(initTime))

	// от количества рабочих в горутинах затрачиваемое время не меняется,
	// т.к. используется работа всего кода "сообща" благодаря горутинам
	// и их параллельной работе
}
