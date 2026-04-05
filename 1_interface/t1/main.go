package main

import "fmt"

type Auto interface {
	StepOnGas()
	StepOnBrake()
}

type BMW struct{}

func (b BMW) StepOnGas() {
	fmt.Println("Я БМВ! 550 лошадиных сил! Жмём на газ!")
}
func (b BMW) StepOnBrake() {
	fmt.Println("Я БМВ! Тормоз в пол! Хорошее торможение!")
}

type Zhiga struct{}

func (z Zhiga) StepOnGas() {
	fmt.Println("Я жига! 2107! Пробую не развалиться!")
}
func (z Zhiga) StepOnBrake() {
	fmt.Println("Я жига! Вот блин тормоза чуть не отвалились(")
}

type Lambo struct{}

func (l Lambo) StepOnGas() {
	fmt.Println("Я ламбоу! 1500 лс! Еду быстро!")
}

/*
func rideBMW(bmw BMW) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю по газам!")
	bmw.StepOnGas()
}

func rideZhiga(zhiga BMW) {
	fmt.Println("Я счастливый обладатель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю по газам!")
	zhiga.StepOnGas()
}
*/

func ride(auto Auto) {
	fmt.Println("Я водитель!")
	fmt.Println("Я сажусь в свою машину!")
	fmt.Println("И нажимаю по газам!")
	auto.StepOnGas()
	auto.StepOnBrake()
}

func main() {
	// bmw := BMW{}
	zhiga := Zhiga{}
	// lambo := Lambo{}

	ride(zhiga)

	// 22:19 2 ролик
}
