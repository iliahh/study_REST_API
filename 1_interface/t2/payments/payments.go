package payments

type PaymentMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

type PaymentModule struct {
	// Как сохранять информацию об оплатах?
	// 1 - статический массив

	// 2 - слайс (норм для малого количества элементов) (если хотим сохранить порядок элементов в памяти)

	// -->> 3 - мапа (кей-велью хранилище)
	// (норм для большого кол-ва элементов)
	// (что-то если хотим сохранить,
	// то оно разбросано в разных ячейках памяти)
	// (мапе порядок элементов не знаком,
	// нельзя отсортировать)

	paymentsInfo map[int]PaymentInfo

	paymentMethod PaymentMethod
}

func NewPaymentModule(paymentMethod PaymentMethod) *PaymentModule {
	return &PaymentModule{
		paymentsInfo:  make(map[int]PaymentInfo),
		paymentMethod: paymentMethod,
	}
}

// Метод Pay()
// Принимает:
// 1. Описание проводимой оплаты
// 2. Сумма оплаты
// Возвращает:
// 1. ID проведённой операции

func (p *PaymentModule) Pay(description string, usd int) int {
	// 1. проводить оплату
	// 2. получать id проведённой оплаты
	id := p.paymentMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		Usd:         usd,
		Cancelled:   false,
	}

	// 3. сохранять информацию о проведённой операции
	// - описание операции
	// - сколько было потрачено
	// - отменённая ли операция
	p.paymentsInfo[id] = info

	// 4. возвращать id проведённой оплаты
	return id
}

// Метод Cancel()
// Принимает:
// 1. ID операции
// Возвращает:
// - ничего не возвращает

func (p *PaymentModule) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentMethod.Cancel(id)

	info.Cancelled = true

	p.paymentsInfo[id] = info
}

// Метод Info()
// Принимает:
// 1. ID операции
// Возвращает:
// Информацию о проведённой операции

func (p *PaymentModule) Info(id int) PaymentInfo {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return PaymentInfo{}
	}

	return info
}

// Метод AllInfo()
// Принимает:
// - ничего не принимает
// Возвращает:
// 1. Информацию о всех проведённых операциях

func (p *PaymentModule) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))

	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}

	return tempMap
}
