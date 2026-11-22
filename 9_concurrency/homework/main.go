package main

import (
	"context"
	"fmt"
	"sync"
	"time"

	"github.com/iliahh/study_REST_API/9_concurrency/homework/sensors/sensPressure"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/sensors/sensSeismic"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/sensors/sensWetness"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/location"
	"github.com/iliahh/study_REST_API/9_concurrency/homework/structs/reading"
	sensordata "github.com/iliahh/study_REST_API/9_concurrency/homework/structs/sensorData"
	"github.com/k0kubun/pp"
)

/*
— Мы пишем симуляцию работы метеоцентра.
— Метеоцентр собирает показания с датчиков трёх видов:
	1. Первый вид датчиков анализирует давление воздуха.
	2. Второй вид датчиков анализирует влажность воздуха.
	3. Третий вид датчиков анализирует сейсмическую активность.
— Необходимо запустить симуляцию каждого вида датчиков в отдельных горутинах,
по одной горутине на датчик. Датчиков каждого вида может быть переменное количество.
— Каждый датчик, помимо прочей информации, передаёт координаты,
откуда передаваемая информация была получена.
— main горутина должна уметь:
	1. Контролировать работу датчиков.
	2. В любомой момент завершать работу одного из видов датчиков.
	3. Собирать всю поступающую от всех датчиков информацию.
— Реализуйте предложенную симуляцию, учитывая все изученные правила
корректной организации работы конкуретного приложения.
*/

func main() { // запускает датчики, собирает и выводит их показания
	mtx := sync.Mutex{} // защищает общие данные от одновременной записи

	// контексты для датчиков трёх типов
	pressureContext, pressureCancel := context.WithCancel(context.Background())
	wetnessContext, wetnessCancel := context.WithCancel(context.Background())
	seismicContext, seismicCancel := context.WithCancel(context.Background())

	go func() { // через три секунды останавливает датчики давления
		time.Sleep(3 * time.Second) // задаёт длительность их работы
		pressureCancel()            // сообщает датчикам, что нужно завершиться
	}()

	go func() {
		time.Sleep(3 * time.Second)
		wetnessCancel()
	}()

	go func() {
		time.Sleep(3 * time.Second)
		seismicCancel()
	}()

	// создаёт две общие координаты для всех датчиков и для сканирования ими координат
	locations := location.GenerateLocations(2)
	// запускают датчики всех типов
	pressureTransferPoint := sensPressure.SensPressurePool(pressureContext, locations)
	wetnessTransferPoint := sensWetness.SensWetnessPool(wetnessContext, locations)
	seismicTransferPoint := sensSeismic.SensSeismicPool(seismicContext, locations)

	// создаёт итоговую запись для каждой координаты
	data := make([]sensordata.SensorData, len(locations))
	for i, loc := range locations { // проходит по всем сгенерированным точкам
		data[i] = sensordata.SensorData{ // сохраняет номер и координаты точки
			SensorID: i + 1, // нумерует точки начиная с единицы
			Location: loc,   // записывает координаты точки
		}
	}

	// ждёт, пока все сборщики прочитают свои каналы
	wg := &sync.WaitGroup{}
	collect := func(ch <-chan reading.Reading) { // читает показания из одного канала
		defer wg.Done() // отмечает завершение сборщика

		for v := range ch { // читает, пока пул не закроет канал
			mtx.Lock()                        // блокирует общие данные перед изменением
			sensorData := &data[v.SensorID-1] // находит итоговую запись по номеру датчика
			switch v.SensorType {             // определяет, какой тип показания пришёл
			case "pressure": // обрабатывает последнее показание давления
				// сохраняет давление
				sensorData.PressureGPa = v.PressureGPa
				// сохраняет время замера
				sensorData.PressureDurationSeconds = v.DurationSeconds
			case "wetness":
				sensorData.WetnessPercent = v.WetnessPercent
				sensorData.WetnessDurationSeconds = v.DurationSeconds
			case "seismic":
				sensorData.SeismicPoint = v.SeismicPoint
				sensorData.SeismicDurationSeconds = v.DurationSeconds
			}
			mtx.Unlock() // разрешает другим сборщикам обновлять данные
		}
	}

	wg.Add(3)                         // учитывает три запускаемых сборщика
	go collect(pressureTransferPoint) // начинает собирать показания давления
	go collect(wetnessTransferPoint)
	go collect(seismicTransferPoint)

	// метеоцентр дождётся закрытия канала, а пул дожидается остановки датчиков
	wg.Wait()
	fmt.Println("--->>  Работа датчиков завершена!")

	fmt.Println("Полученные данные:")
	mtx.Lock()       // защищает итоговый вывод от параллельного изменения данных
	pp.Println(data) // выводит объединённые показания
	mtx.Unlock()     // освобождает мьютекс после вывода

	/*
			Полученные данные:
		[]main.SensorData{
		  main.SensorData{
		    SensorID: 1,
		    Location: location.Location{
		      Latitude:  78.782313,
		      Longitude: 137.704651,
		    },
		    PressureGPa:             1000.076232,
		    WetnessPercent:          82,
		    SeismicPoint:            5,
		    PressureDurationSeconds: 0.378194,
		    WetnessDurationSeconds:  0.557341,
		    SeismicDurationSeconds:  0.607644,
		  },
		  main.SensorData{
		    SensorID: 2,
		    Location: location.Location{
		      Latitude:  20.269504,
		      Longitude: -25.109830,
		    },
		    PressureGPa:             985.278628,
		    WetnessPercent:          89,
		    SeismicPoint:            10,
		    PressureDurationSeconds: 0.834937,
		    WetnessDurationSeconds:  0.514733,
		    SeismicDurationSeconds:  0.187136,
		  },
		}
	*/
}
