package payments

type paymentsMethod interface{

	pay(usd int)int
    cancell(id int)
}

type Paymentst struct{

paymentsStruct   PaymentsStruct

}
func NewpaymentModul(   paymentsmethod  paymentsMethod )*PaymentModul{
	return &PaymentModul{
paymentsmethod:paymentsmethod,


}
}





func ( p PaymentModul) Pay(destriction string,usd int)int {

id:=paymentsModul.Pay(usd)int


info:=Destriction  desrtiction
      usd          Usd
	  cancell      false

}





func (p PaymentModul) Cancell(){}

func (p PaymentModul) Info(){}

func (p PaymentModul) Allinfo(){}
