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
	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		box["age"] = 25
		box["name"] = "Ilya"
		box["active"] = true
		mtx.Unlock()
	}
}

func data2(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		box["age"] = 23
		box["name"] = "Senya"
		box["active"] = false
		mtx.Unlock()
	}
}

// горутина-читатель
func readData(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		_ = box["age"]
		mtx.Unlock()
	}
}

func main() {
	// - мы с вами пишем универсальное хранилище, основанное на мапе
	uniBox := make(map[string]any)
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go data1(wg, uniBox)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go data2(wg, uniBox)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go readData(wg, uniBox)
	}

	wg.Wait() // дождаться всех горутин-писателей и -читателей

	fmt.Println(uniBox)

	elapsed := time.Since(initTime)
	fmt.Printf("Время выполнения: %.9f сек.\n", elapsed.Seconds())

	// map[active:false age:23 name:Senya]
	// Время выполнения: 0.128038800 сек.

	// -> go run -race main.go
	// map[active:true age:25 name:Ilya]
	// Время выполнения: 0.914273000 сек.
}
*/

// защита от гонки данных теперь через рвмутекс
var mtx sync.RWMutex

// поэтому записывать будем сразу из горутины данные в мапу
func data1(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	// - ключом в хранилище является строка (строки-ключи ниже в кавычках)
	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		// defer mtx.Unlock()

		box["age"] = 25
		box["name"] = "Ilya"
		box["active"] = true
		mtx.Unlock()
	}
}

func data2(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	for i := 1; i <= 100_000; i++ {
		mtx.Lock()
		// defer mtx.Unlock()

		box["age"] = 23
		box["name"] = "Senya"
		box["active"] = false
		mtx.Unlock()
	}
}

// горутина-читатель
func readData(wg *sync.WaitGroup, box map[string]any) {
	defer wg.Done()

	// для записи используются Lock/Unlock для записи,
	// для чтения используется mtx.Rlock()/mtx.RUnlock()
	for i := 1; i <= 100_000; i++ {
		mtx.RLock()
		_ = box["age"]
		mtx.RUnlock()
	}
}

func main() {
	// - мы с вами пишем универсальное хранилище, основанное на мапе
	uniBox := make(map[string]any)
	wg := &sync.WaitGroup{}

	initTime := time.Now()

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go data1(wg, uniBox)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go data2(wg, uniBox)
	}

	for i := 1; i <= 10; i++ {
		wg.Add(1)
		go readData(wg, uniBox)
	}

	wg.Wait() // дождаться всех горутин-писателей и -читателей

	fmt.Println(uniBox)

	elapsed := time.Since(initTime)
	fmt.Printf("Время выполнения: %.9f сек.\n", elapsed.Seconds())

	// map[active:true age:25 name:Ilya]
	// Время выполнения: 0.209301800 сек.

	// -> go run -race main.go
	// map[active:true age:25 name:Ilya]
	// Время выполнения: 0.849453800 сек.
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

// в моем случае, замеряя время выполнения кода при использовании разных
// примитивов синхронизации, Мьютекс оказался быстрее, чем РВМьютекс
