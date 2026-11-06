package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

/*
- опишите функцию, которая имитирует огородника, который удобряет
и поливает определённую область огромного ботанического сада.
- каждый огородник удобряет и поливает свою область какое-то случайное
количество времени (500ms - 1000ms).
- при этом количество огородников тоже случайное, ведь какая-то почва ещё могла не просохнуть
до конца и не нуждаться в поливе на настоящий момент.
- запустите в отдельных горутинах случайное количество огородников, дождитесь выполнения каждого
(никаких time.Sleep для ожидания огородников), после чего в main горутине выведите на экран
информацию об окончании выполнения программы.
*/

/*
- опишите функцию, которая имитирует огородника, который удобряет
и поливает определённую область огромного ботанического сада.
*/
func ogorodnik(wg *sync.WaitGroup, area string) {
	defer wg.Done() // гарантированная отработка при завершении горутины

	fmt.Printf("Огородник удобряет и поливает %s область сада\n", area)

	// какое-то случайное количество времени (500ms - 1000ms)
	worktime := time.Duration(500+rand.Intn(501)) * time.Millisecond

	time.Sleep(worktime)

	fmt.Printf("Область сада %s обработана\n", area)
}

func main() {
	wg := &sync.WaitGroup{}

	// случайное количество огородников
	for i := 0; i < rand.Intn(10)+1; i++ {
		wg.Add(1) // до запуска горутины
		go ogorodnik(wg, fmt.Sprintf("%d", i+1))
	}

	wg.Wait()

	fmt.Println("Все области ботанического сада удобрены и политы")
}

// мб через мапы (ключи) задать бронь определенной области за определенным огородником,
// так еще и рандомом?
/*
func ogorodnik(wg *sync.WaitGroup, name, area string) {
    defer wg.Done()

    fmt.Printf("%s удобряет и поливает участок %s\n", name, area)

    worktime := time.Duration(500+rand.Intn(501)) * time.Millisecond
    time.Sleep(worktime)

    fmt.Printf("%s закончил работу на участке %s\n", name, area)
}

func main() {
    areas := []string{"северный", "южный", "восточный", "западный"}

    // Перемешиваем участки для случайного назначения
    rand.Shuffle(len(areas), func(i, j int) {
        areas[i], areas[j] = areas[j], areas[i]
    })

    count := rand.Intn(len(areas)) + 1
    assignments := make(map[string]string, count)

    for i := 0; i < count; i++ {
        name := fmt.Sprintf("Огородник-%d", i+1)
        assignments[name] = areas[i]
    }

    var wg sync.WaitGroup
    wg.Add(len(assignments))

    for name, area := range assignments {
        go ogorodnik(&wg, name, area)
    }

    wg.Wait()
    fmt.Println("Все области ботанического сада удобрены и политы")
}
*/
