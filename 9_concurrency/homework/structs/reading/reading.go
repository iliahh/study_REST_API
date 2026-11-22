package reading

import "github.com/iliahh/study_REST_API/9_concurrency/homework/structs/location"

/*
	объединяет собираемые данные с номерами и координатами датчиков:
	одно показание одного типа - давление, влажность или активность
*/
type Reading struct {
	SensorID        int
	SensorType      string
	Location        location.Location
	PressureGPa     float64
	WetnessPercent  int
	SeismicPoint    int
	DurationSeconds float64
}
