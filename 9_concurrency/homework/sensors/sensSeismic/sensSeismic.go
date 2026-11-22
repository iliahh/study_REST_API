package sensSeismic

import (
	"context"
	"fmt"
	"math/rand/v2"
	"sync"
	"time"

	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/location"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/reading"
)

func SensSeismicPool(ctx context.Context,
	locations []location.Location,
) <-chan reading.Reading {
	// канал передаёт показания датчиков
	seismicTransferPoint := make(chan reading.Reading)
	wg := &sync.WaitGroup{}

	for i, location := range locations {
		wg.Add(1)
		go seismicSensor(ctx, wg, seismicTransferPoint, i+1, location)
	}

	go func() {
		wg.Wait()
		close(seismicTransferPoint)
	}()

	return seismicTransferPoint
}

func seismicSensor(ctx context.Context,
	wg *sync.WaitGroup,
	transferPoint chan<- reading.Reading,
	sensorID int,
	location location.Location,
) {
	defer wg.Done() // датчик завершил работуwetnessSensor

	ticker := time.NewTicker(time.Second)
	defer ticker.Stop()

	fmt.Printf("Датчик сейсмической активности №%d начал сбор данных\n", sensorID)

	sendReading := func() {
		reading := reading.Reading{
			SensorID:        sensorID,
			SensorType:      "seismic",
			Location:        location,
			SeismicPoint:    rand.IntN(12) + 1,        // активность от 1 до 12 баллов
			DurationSeconds: 0.1 + rand.Float64()*0.9, // время в секундах
		}
		transferPoint <- reading // передаем в канал данные формата структуры ридинг
		fmt.Printf("Датчик сейсмической активности №%d передал собранные данные\n", sensorID)
	}

	for {
		select {
		case <-ctx.Done(): // контекст отменён - датчик прекращает измерения
			fmt.Printf("Датчик сейсмической активности №%d закончил сбор данных\n", sensorID)
			return

		case <-ticker.C:
			// здесь генерируется условная влажность в диапазоне 20-90 %
			fmt.Printf("Датчик сейсмической активности №%d получил данные\n", sensorID)
			sendReading()
		}
	}
}
