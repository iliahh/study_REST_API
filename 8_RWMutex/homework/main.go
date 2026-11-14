package main

import (
	"fmt"
	"sync"
	"time"
)

/*
- мы с вами пишем универсальное хранилище, основанное на мапе
- ключом в хранилище является строка
- необходимо реализовать это хранилище, допуская вероятность конкурентного использования
- напишите нагрузочные тесты (такие, как мы писали в уроке в теме RWMutex),
и сравните время выполнения программы с использованием Mutex и с использованием RWMutex
(p.s. не всегда будет прирост в скорости, не нужно этого пугаться)
*/

/*
// защищаться от гонки данных будем мутексом
var mtx sync.Mutex

// поэтому записывать будем сразу из горутины данные в мапу
func data1(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	// - ключом в хранилище является строка (строки-ключи ниже в кавычках)
	mtx.Lock()
	defer mtx.Unlock()
	box["age"] = 25
	box["name"] = "Ilya"
	box["active"] = true
}

func data2(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	mtx.Lock()
	defer mtx.Unlock()
	box["age"] = 23
	box["name"] = "Senya"
	box["active"] = false
}

func main() {
	// - мы с вами пишем универсальное хранилище, основанное на мапе
	uniBox := make(map[string]any)
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	wg.Add(2)
	go data1(wg, uniBox)
	go data2(wg, uniBox)

	wg.Wait()

	fmt.Println(uniBox)

	elapsed := time.Since(initTime)
	fmt.Printf("Время выполнения: %.9f сек.\n", elapsed.Seconds())

	// -> go run -race main.go
	// map[active:false age:23 name:Senya]
	// Время выполнения: 0.000515300 сек.
}
*/

// защита от гонки данных теперь через рвмутекс
var mtx sync.RWMutex

// поэтому записывать будем сразу из горутины данные в мапу
func data1(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	// - ключом в хранилище является строка (строки-ключи ниже в кавычках)
	mtx.Lock()
	defer mtx.Unlock()
	box["age"] = 25
	box["name"] = "Ilya"
	box["active"] = true
}

func data2(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	mtx.Lock()
	defer mtx.Unlock()
	box["age"] = 23
	box["name"] = "Senya"
	box["active"] = false
}

func main() {
	// - мы с вами пишем универсальное хранилище, основанное на мапе
	uniBox := make(map[string]any)
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	wg.Add(2)
	go data1(wg, uniBox)
	go data2(wg, uniBox)

	wg.Wait()

	// для записи используются Lock/Unlock для записи,
	// для чтения используется mtx.Rlock()/mtx.RUnlock()
	mtx.RLock()
	fmt.Println(uniBox)
	mtx.RUnlock()

	elapsed := time.Since(initTime)
	fmt.Printf("Время выполнения: %.9f сек.\n", elapsed.Seconds())

	// -> go run -race main.go
	// map[active:false age:23 name:Senya]
	// Время выполнения: 0.000515700 сек.
}

/*
	В примере ставится на запись Lock/Unlock, на чтение RLock/RUnlock:

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
*/
