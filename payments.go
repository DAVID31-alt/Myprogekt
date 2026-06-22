package payments

type paymentsMethod interface{




	pay(usd int)int
    cancell(id int)
}


p.paymentsInfo map[int]PaymentInfo


paymentsmethod paymentsMethod

func NewpaymentModul(   paymentsmethod  paymentsMethod )*PaymentModul{
	return &PaymentModul{
		paymentsInfo make(map[int]PaymentInfo ,)
paymentsmethod:paymentsmethod,


}
}





func ( p PaymentModul) Pay(destriction string,usd int)int {

id:=p.paymentsMethod.Pay(usd)
}

info:=PaymentInfo{
      Destriction  desrtiction
      usd          Usd
	  cancelled      false

}

p.paymentsInfo[id]=info
return id



func (p PaymentModul) Cancell(id int){
info,ok:=p.paymetsInfo[id]

if !ok{
	return
}

p.paymentsMethod.cancell(id)
info.cancelled=true
p.paymentsInfo[id]=info
}


func (p PaymentModul) Info(id int){}

info,ok:=p.paymentsInfo[id]
if !ok{
return paymentsInfo{}
}


func (p PaymentModul) Allinfo map[int]PaymentsInfo{
	return p.paymentsInfo
}
