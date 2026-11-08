package main

import (
	"math/rand"
	"sync"

	"github.com/k0kubun/pp"
)

/*
- представьте что вы - знаменитость.
- фанаты, независимо друг от друга, кладут очень много писем вам в почтовый ящик.
- создайте слайс строк, где каждая строка - это письмо от фаната.
- каждый фанат - это отдельная горутина, которая много раз пишет вам новое письмо
(добавляет новый элемент в слайс).
- запустите программу и оцените, есть ли в представленной задаче гонка данных или нет
- как можно отловить гонку данных?
- если гонка данных есть, избавьтесь от неё, после чего проведите повторные проверки.
*/

/*
// горутина фаната и мейн горутина, если бы я делал гонку данных:
func fan(id int, wg *sync.WaitGroup, mailbox *[]Letter) {
    defer wg.Done()

    numMail := rand.Intn(5) + 1
    for range numMail {
        *mailbox = append(*mailbox, Letter{
            Fan:  id,
            Mail: mail(),
        })
    }
}

func main() {
    mailbox := make([]Letter, 0)
    var wg sync.WaitGroup

    wg.Add(3)
    go fan(1, &wg, &mailbox)
    go fan(2, &wg, &mailbox)
    go fan(3, &wg, &mailbox)

    wg.Wait()
    pp.Println(mailbox)
}

// гонки данных в моем случае нет, потому что я отправляю фанатами письма в канал,
// и добавляет эти письма в слайс уже мейн, так доступ в мейлбокс безопасен от гонки данных

*/

type Letter struct {
	Fan  int
	Mail string
}

func mail() string {
	// генерируем случайное письмо фаната из 5-ти символов
	mail := make([]byte, 5) // создаём слайс из 5 байтов, место для хранения пяти символов

	for i := range mail { // 0, 1, 2, 3, 4
		mail[i] = byte(rand.Intn(26) + 97) // генерируем случайный символ от 'a' до 'z'
	}

	return string(mail) // преобразовываем слайс байтов в строчный слайс и возвращаем его
}

// func fan1(wg *sync.WaitGroup, ch chan string)
func fan1(id int, wg *sync.WaitGroup, ch chan Letter) {
	defer wg.Done()

	numMail := rand.Intn(5) + 1
	for range numMail {
		ch <- Letter{
			Fan:  id,
			Mail: mail(),
		}
	}
}

func fan2(id int, wg *sync.WaitGroup, ch chan Letter) {
	defer wg.Done()

	numMail := rand.Intn(5) + 1
	for range numMail {
		ch <- Letter{
			Fan:  id,
			Mail: mail(),
		}
	}
}

func fan3(id int, wg *sync.WaitGroup, ch chan Letter) {
	defer wg.Done()

	numMail := rand.Intn(5) + 1
	for range numMail {
		ch <- Letter{
			Fan:  id,
			Mail: mail(),
		}
	}
}

func main() {
	ch := make(chan Letter) // ch := make(chan string)
	wg := &sync.WaitGroup{}

	mailbox := make([]Letter, 0) // mailbox := make([]string, 0) - почтовый ящик для писем

	wg.Add(3)
	go fan1(1, wg, ch)
	go fan2(2, wg, ch)
	go fan3(3, wg, ch)

	go func() {
		wg.Wait()
		close(ch)
	}()

	for letter := range ch {
		mailbox = append(mailbox, letter)
	}

	pp.Println(mailbox)
}
