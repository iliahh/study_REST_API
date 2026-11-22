package sensPressure

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"

	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/location"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/reading"
)

// запускает по одной горутине на каждую переданную координату и возвращает канал,
// из которого можно читать показания всех данных с датчиков
func SensPressurePool(ctx context.Context,
	locations []location.Location,
) <-chan reading.Reading {
	// канал передаёт показания датчиков
	pressureTransferPoint := make(chan reading.Reading)
	wg := &sync.WaitGroup{}

	for i, location := range locations {
		wg.Add(1) // учитываем новую горутину, чтобы дождаться её завершения.
		go pressureSensor(ctx, wg, pressureTransferPoint, i+1, location)
	}

	// пока читается канал с отправленными данными от горутины-датчика,
	// то канал не закроется, но как горутина-датчик завершится, то и закроется канал
	go func() {
		wg.Wait()
		close(pressureTransferPoint)
	}()

	return pressureTransferPoint
}

func pressureSensor(
	ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- reading.Reading,
	sensorID int,
	location location.Location,
) {
	defer wg.Done() // датчик завершил работу

	/*
		time.NewTicker - функция стандартного пакета time.
		она возвращает объект Ticker, у которого есть поле-канал C.
		имя C выбрано авторами пакета, это обычное поле структуры.

		тикер нужен, когда действие должно повторяться с интервалом.
		для одноразовой задержки подойдёт, например, time.Sleep,
		а если повторять действие не нужно - тикер вообще не нужен.
		если создал тикер, его обычно останавливают через defer ticker.Stop().
	*/
	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	fmt.Printf("Датчик давления №%d начал сбор данных\n", sensorID)

	sendReading := func() {
		reading := reading.Reading{
			SensorID:        sensorID,
			SensorType:      "pressure",
			Location:        location,
			PressureGPa:     980 + rand.Float64()*70,  // давление в гПа
			DurationSeconds: 0.1 + rand.Float64()*0.9, // время в секундах
		}
		transferPoint <- reading // передаем в канал данные формата структуры ридинг
		fmt.Printf("Датчик давления №%d передал собранные данные\n", sensorID)
	}

	for {
		select {
		case <-ctx.Done(): // контекст отменён - датчик прекращает измерения
			fmt.Printf("Датчик давления №%d закончил сбор данных\n", sensorID)
			return

		case <-ticker.C:
			// здесь генерируется условное давление в диапазоне 980-1050 гПа
			fmt.Printf("Датчик давления №%d получил данные\n", sensorID)
			sendReading()
		}
	}
}
