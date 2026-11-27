package main

import (
	"fmt"
	"net/http"
	"time"
)

// если написать func handler(wr), то вс код сразу подставит
// снипет с РеспонсВрайтер и Реквест

// программа, слушающая входящие сетевые подключения на 9091 порту
// и предоставляющая три различных хендлера по трём различным паттернам:

/*
func payHandler(w http.ResponseWriter, r *http.Request) {
	str := "Новый платёж обработан!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно совершил оплату!")
	}
}

func cancelHandler(w http.ResponseWriter, r *http.Request) {
	str := "Оплата отменена!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно отменил оплату!")
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	str := "Hello, world!"
	b := []byte(str)
	_, err := w.Write(b)
	if err != nil {
		fmt.Println("Во время записи HTTP ответа произошла ошибка:", err.Error())
	} else {
		fmt.Println("Я корректно обработал HTTP запрос!")
	}
}

func main() {
	// http.HandleFunc(pattern string, handler func(http.ResponseWriterm *http.Request))
	http.HandleFunc("/default", handler)
	http.HandleFunc("/pay", payHandler)
	http.HandleFunc("/cancel", cancelHandler)

	// запускаем сервер и блокируем на этой строке выполнение нашего кода
	// и слушает входящие сетевые подключения на наши паттерны,
	// а далее вызывает наши хендлеры
	fmt.Println("Запускаю HTTP сервер!")
	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("Произошла ошибка:", err.Error())
	}

	fmt.Println("Программа закончила своё выполнение!")
}
*/

// под каждым хендлером под капотом запускается своя горутина!
func handler(w http.ResponseWriter, r *http.Request) {
	_, err := w.Write([]byte("Hello, world!"))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}

	fmt.Println("Handler отработал успешно!")
}

func handlerSleep(w http.ResponseWriter, r *http.Request) {
	time.Sleep(5 * time.Second)

	_, err := w.Write([]byte("HTTP response!"))
	if err != nil {
		fmt.Println("err:", err.Error())
		return
	}
}

func main() {
	// "/" - эндпоинтами называют еще вместо паттернов
	// 	  youtube.com/
	http.HandleFunc("/", handler)
	// localhost:9091 -> запросы просто по домену + порт,
	// но сам паттерн указывать не нужно,
	// автоматически будет вызван пустой паттерн

	http.HandleFunc("/sleep", handlerSleep)

	err := http.ListenAndServe(":9091", nil)
	if err != nil {
		fmt.Println("err:", err.Error())
	}
}
