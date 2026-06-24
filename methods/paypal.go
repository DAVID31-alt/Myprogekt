package methods

import (
	"fmt"
	"math/rand"
)

type paypal struct{}

func (c paypal) pay(usd int) int {

	fmt.Println("оплата пей палом")
	fmt.Println("Размер оплаты", usd, "USDT")

	id := rand.Int()

	return id

}
func (c paypal) Cancel(id int) {
	fmt.Println("отмена оплаты paypal! Id", id)
}
