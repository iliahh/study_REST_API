package main

import (
	"fmt"
	"sync"
	"time"
)

var likes int = 0

// Позволяет горутинкам читать данные параллельно, не блокируя друг друга, но при этом блокирует запись данных.
// Если мы видим, что нам не нужно использовать блокировки на чтение, то лучше использовать обычный мьютекс, так как он быстрее,
// в ином случае, если у нас много горутин, которые читают данные, то лучше использовать RWMutex.
var mtx sync.RWMutex

func setLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.Lock() // блокировка на запись
		likes++
		mtx.Unlock() // снятие блокировки
	}
}

func getLike(wg *sync.WaitGroup) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.RLock() // блокировка на чтение
		_ = likes
		mtx.RUnlock() // снятие блокировки
	}
}

func main() {
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go setLike(wg)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go getLike(wg)
	}

	wg.Wait()

	fmt.Println("Время выполнения:", time.Since(initTime))
}
