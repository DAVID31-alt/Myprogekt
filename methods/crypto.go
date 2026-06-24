package methods

import (
	"fmt"
	"math/rand"
)

type Crypto struct{}

func (c Crypto) pay(usd int) int {

	fmt.Println("оплата криптоволютой")
	fmt.Println("Размер оплаты", usd, "USDT")

	id := rand.Int()

	return id

}
func (c Crypto) Cancel(id int) {
	fmt.Println("отмена крипто операции! Id", id)
}
