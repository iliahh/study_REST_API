package postman

import (
	"context"
	"fmt"
	"sync"
	"time"
)

func Postman(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- string,
	n int, mail string,
) {
	// после запуска каждого почтальона,
	// в конце его работы мы будем сигнализировать о том,
	// что его работа завершилась
	defer wg.Done()

	for {
		select {
		case <-ctx.Done():
			fmt.Printf("Я почтальон №%d, мой рабочий день закончен!\n", n)
			return
		default: // секция в дефаулте выполняется, пока контекст еще не отменён
			fmt.Printf("Я почтальон №%d, взял письмо\n", n)
			time.Sleep(1 * time.Second) // взяли и несём письмо
			fmt.Printf("Я почтальон №%d, донёс письмо до почты: %s\n", n, mail)

			transferPoint <- mail
			fmt.Printf("Я почтальон №%d, передал письмо : %s\n", n, mail)
		}
	}
}

// функция, принимающая количество почтальонов и запускающая сразу
// несколько их горутин в работу
func PostmanPool(ctx context.Context, postmanCount int) <-chan string {
	// инициализируем канал, который требуют почтальоны для передачи писем
	mailTransferPoint := make(chan string)

	wg := &sync.WaitGroup{}

	for i := 1; i <= postmanCount; i++ {
		wg.Add(1)
		go Postman(ctx, wg, mailTransferPoint, i, postmanToMail(i))
	}

	go func() {
		wg.Wait()
		close(mailTransferPoint)
	}()

	return mailTransferPoint
}

func postmanToMail(postmanNumber int) string {
	ptm := map[int]string{
		1: "Семейный привет",
		2: "Приглашение от друга",
		3: "Информация из автосервиса",
	}

	mail, ok := ptm[postmanNumber]
	if !ok {
		return "Лотерея"
	}

	return mail
}
