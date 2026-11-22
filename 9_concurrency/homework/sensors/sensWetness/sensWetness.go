package sensWetness

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/location"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/reading"
)

func SensWetnessPool(ctx context.Context,
	locations []location.Location,
) <-chan reading.Reading {
	// канал передаёт показания датчиков
	wetnessTransferPoint := make(chan reading.Reading)
	wg := &sync.WaitGroup{}

	for i, location := range locations {
		wg.Add(1)
		go wetnessSensor(ctx, wg, wetnessTransferPoint, i+1, location)
	}

	go func() {
		wg.Wait()
		close(wetnessTransferPoint)
	}()

	return wetnessTransferPoint
}

func wetnessSensor(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- reading.Reading,
	sensorID int,
	location location.Location,
) {
	defer wg.Done() // датчик завершил работуwetnessSensor

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	fmt.Printf("Датчик влажности №%d начал сбор данных\n", sensorID)

	sendReading := func() {
		reading := reading.Reading{
			SensorID:        sensorID,
			SensorType:      "wetness",
			Location:        location,
			WetnessPercent:  rand.IntN(71) + 20,       // влажность от 20 до 90 %
			DurationSeconds: 0.1 + rand.Float64()*0.9, // время в секундах
		}
		transferPoint <- reading // передаем в канал данные формата структуры ридинг
		fmt.Printf("Датчик влажности №%d передал собранные данные\n", sensorID)
	}

	for {
		select {
		case <-ctx.Done(): // контекст отменён - датчик прекращает измерения
			fmt.Printf("Датчик влажности №%d закончил сбор данных\n", sensorID)
			return

		case <-ticker.C:
			// здесь генерируется условная влажность в диапазоне 20-90 %
			fmt.Printf("Датчик влажности №%d получил данные\n", sensorID)
			sendReading()
		}
	}
}
