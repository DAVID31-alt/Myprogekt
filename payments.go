package main

// PaymentInfo хранит данные о конкретном платеже
type PaymentInfo struct {
	Description string
	USD         int
	Cancelled   bool
}

// PaymentsMethod описывает интерфейс платежной системы
type PaymentsMethod interface {
	Pay(usd int) int
	Cancel(id int)
}

// PaymentModul управляет всеми платежами
type PaymentModul struct {
	paymentsInfo   map[int]PaymentInfo
	paymentsMethod PaymentsMethod
}

// NewPaymentModul — конструктор модуля
func NewPaymentModul(method PaymentsMethod) *PaymentModul {
	return &PaymentModul{
		paymentsInfo:   make(map[int]PaymentInfo),
		paymentsMethod: method,
	}
}

// Pay совершает платеж и сохраняет информацию о нем
func (p *PaymentModul) Pay(description string, usd int) int {
	id := p.paymentsMethod.Pay(usd)

	info := PaymentInfo{
		Description: description,
		USD:         usd,
		Cancelled:   false,
	}

	p.paymentsInfo[id] = info
	return id
}

// Cancel отменяет платеж по ID
func (p *PaymentModul) Cancel(id int) {
	info, ok := p.paymentsInfo[id]
	if !ok {
		return
	}

	p.paymentsMethod.Cancel(id)
	info.Cancelled = true
	p.paymentsInfo[id] = info
}

// Info возвращает информацию о конкретном платеже
func (p *PaymentModul) Info(id int) (PaymentInfo, bool) {
	info, ok := p.paymentsInfo[id]
	return info, ok
}

// AllInfo возвращает глубокую копию всей карты платежей
func (p *PaymentModul) AllInfo() map[int]PaymentInfo {
	tempMap := make(map[int]PaymentInfo, len(p.paymentsInfo))
	for k, v := range p.paymentsInfo {
		tempMap[k] = v
	}
	return tempMap
}
