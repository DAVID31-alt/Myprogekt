package main

type PaymentsMethod interface {
	pay(usd int) int

	cancell(id int)
}

type PaymentsModul struct{


 
infomap	map[int]paymentsInfo

paymentsMethod: PaymentsMethod
}






func newPAymentsModul(paymentsMethod PaymentsMethod) *PaymentsModul {
	return &PaymentsModul{
PaymentsInfo:=make(map[int]paymentsInfo, )
		paymentsMethod: paymentsMethod,
	}
}

func (p PaymentsModul) pay(Description string, usd int) int {
	id := p.PaymentsMethod.pay(usd)

	info := paymentsInfo{
		Description: description,

		Usd:         usd,

		cancell:     false,
	}
}
 

p.infomap[id]=info{

	return id
}

type paymentsInfo struct {
	Description string

	Usd int

	cancell boll
}

func (p PaymentsModul) cancell(id int) {}

info,ok:=p.paymentsInfo[id]

if !ok{

	return
}

info.cancell=true


p.PaymentsMethod.cancell[id]


p.paymentInfo[id]=info


func (p PaymentsModul) info(id int) PaymentInfo{

info,ok:=p.paymentsInfo[id]
if !ok{
	return PaymentsInfo{}
}
}


func (p PaymentsModul) allinfo()map[int]PaymentsInfo {

tempmap:= make(map[int]PaymentsInfo,len(p.paymentsInfo))

for k,v:=range p.paymentsInfo{

	tempMap[k]=v
}


	return p.paymentsInfo



}

