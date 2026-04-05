package main

import (
	"study/payments"
	"study/payments/methods"

	"github.com/k0kubun/pp"
)

/*
// slice vs map

// slice:
// O(n)
// 1 -> 10 us
// 10 ->> 100 us
// 100 -> 1000 us
// 1_000_000 -> 10_000_000 -> 10 s

// map:
// 0(1)
// 1 -> 10 us
// 10 -> 10 us
// 100 -> 10 us
// 1_000_000 -> 10 us

// 1. сохранять большое количество PaymentInfo
// 2. находить конкретный элемент PaymentInfo в множестве

type PaymentInfo struct {
	ID          int
	Description string
	Usd         int
	Cancelled   bool
}

type PaymentModuleWithSlice struct {
	s []PaymentInfo
}

func (m *PaymentModuleWithSlice) AddInfo(info PaymentInfo) {
	m.s = append(m.s, info)
}

func (m *PaymentModuleWithSlice) GetInfo(id int) PaymentInfo {
	for _, info := range m.s {
		if info.ID == id {
			return info
		}
	}

	return PaymentInfo{}
}

// -----------------------------------------------------

type PaymentModuleWithMap struct {
	m map[int]PaymentInfo
}

func (m *PaymentModuleWithMap) AddInfo(info PaymentInfo) {
	m.m[info.ID] = info
}

func (m *PaymentModuleWithMap) GetInfo(id int) PaymentInfo {
	info, ok := m.m[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

func main() {
	pSlice := PaymentModuleWithSlice{}
	pMap := PaymentModuleWithMap{
		m: make(map[int]PaymentInfo),
	}

	interations := 100_000

	// ---
	before := time.Now()

	for i := 0; i < interations; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}

		pSlice.AddInfo(info)
	}

	fmt.Println("slice add:", time.Since(before))

	// Тоже самое, но для мапы:
	before = time.Now()

	for i := 0; i < interations; i++ {
		info := PaymentInfo{
			ID:          i,
			Description: strconv.Itoa(i),
		}

		pMap.AddInfo(info)
	}

	fmt.Println("map add:", time.Since(before))
	// -----------------------------------------------------

	before = time.Now()

	for i := 0; i < interations; i++ {
		info := pSlice.GetInfo(i)
		_ = info
	}

	fmt.Println("slice get:", time.Since(before))

	before = time.Now()

	for i := 0; i < interations; i++ {
		info := pMap.GetInfo(i)
		_ = info
	}

	fmt.Println("map get:", time.Since(before))

	// info1 := PaymentInfo{
	// 	ID:          10,
	// 	Description: "BURGER",
	// 	Usd:         5,
	// 	Cancelled:   false,
	// }

	// pSlice.AddInfo(info1)
	// pMap.AddInfo(info1)

	// // pp.Println("pSlice:", pSlice)
	// // pp.Println("pMap:", pMap)

	// i1 := pSlice.GetIngo(10)
	// i2 := pMap.GetInfo(10)

	// fmt.Println("i1:", i1)
	// fmt.Println("i2:", i2)
}
*/

func main() {
	method := methods.NewCrypto()
	// method := methods.NewBank() - смена потока финансовых операций

	paymentModule := payments.NewPaymentModule(method)

	paymentModule.Pay("Бургер", 5)
	idPhone := paymentModule.Pay("Телефон", 500)
	idGame := paymentModule.Pay("Игра", 20)

	paymentModule.Cancel(idPhone)

	allInfo := paymentModule.AllInfo()

	pp.Println("Все наши оплаты:", allInfo)

	gameInfo := paymentModule.Info(idGame)
	pp.Println("Game info:", gameInfo)
}
