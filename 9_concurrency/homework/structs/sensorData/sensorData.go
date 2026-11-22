package sensordata

import "github.com/iliahh/study_REST_API/9_concurrency/homework/structs/location"

// объединяет последние показания всех типов для одной координаты в одну запись
type SensorData struct {
	SensorID                int
	Location                location.Location
	PressureGPa             float64
	WetnessPercent          int
	SeismicPoint            int
	PressureDurationSeconds float64
	WetnessDurationSeconds  float64
	SeismicDurationSeconds  float64
}
