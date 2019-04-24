package entity

import (
	"github.com/dirkarnez/go-dynamic-proxy/pogo"
	"github.com/jinzhu/copier"
)

type entityFactory struct {

}

var (
	Factory = entityFactory{}
)


func (e *entityFactory) NewBill() IBill {
	return &bill{ before: nil, Bill: &pogo.Bill{}}
}

type bill struct {
	before *pogo.Bill
	*pogo.Bill
}

type IBill interface {
	SetPrice(int)
	Change() (*pogo.Bill, *pogo.Bill)
	GetPure() *pogo.Bill
	StartAudit()
}

func (b *bill) Change() (*pogo.Bill, *pogo.Bill) {
	return b.before, b.Bill
}

func (b *bill) SetPrice(price int)  {
	b.Price = price
}

func (b *bill) GetPure() *pogo.Bill {
	return b.Bill
}

func (b *bill) StartAudit() {
	var old *pogo.Bill
	if b.ID == 0 {
		old = nil
	} else {
		old = &pogo.Bill{}
		copier.Copy(old, b)
	}
	b.before = old
}