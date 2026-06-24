package methods

import (
	"fmt"
	"math/rand"
)

type cart struct{}

func (cart) pay(usd int) int {

	fmt.Println("оплата криптоволютой")
	fmt.Println("Размер оплаты", usd, "USDT")

	id := rand.Int()

	return id

}
func (c cart) Cancel(id int) {
	fmt.Println("отмена операци банковской картой  Id", id)
}
