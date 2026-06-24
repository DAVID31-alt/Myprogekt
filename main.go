package main

import "fmt"

// 1. Интерфейс рынка (правила для внешнего магазина)
type garagemethod interface {
	buy(usd int) int
	sell(model string) int
}

// 2. Структура машины
type car struct {
	destriction string
	price       int
	status      bool
}

// 3. Главная структура Тюнинг-Клуба
type garageModul struct {
	mapinfo map[int]car
	nextID  int
	market  garagemethod
}

// 4. Функция-конструктор для включения гаража
func NewGargeModul(market garagemethod) *garageModul {
	return &garageModul{
		mapinfo: make(map[int]car),
		market:  market,
		nextID:  1,
	}
}

// 5. Метод покупки машины в клуб
func (g *garageModul) buy(destriction string, usd int) int {
	_ = g.market.buy(usd) // Симулируем покупку на внешнем рынке

	myCarID := g.nextID
	g.nextID++

	newCar := car{
		destriction: destriction,
		price:       usd,
		status:      false,
	}

	g.mapinfo[myCarID] = newCar
	return myCarID
}

// 6. Метод продажи машины и удаления её из клуба
func (g *garageModul) sell(id int) int {
	mycar := g.mapinfo[id]
	money := g.market.sell(mycar.destriction)

	delete(g.mapinfo, id)
	return money
}

// 7. Метод просмотра всех машин клуба
func (g garageModul) allinfocar() map[int]car {
	return g.mapinfo
}

// 8. Метод подсчета общего количества машин в клубе
func (g garageModul) infogarage() int {
	return len(g.mapinfo)
}

func main() {
	fmt.Println("Тюнинг-клуб успешно спроектирован!")
}
