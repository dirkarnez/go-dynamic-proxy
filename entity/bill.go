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
	GetPtr() *pogo.Bill
	StartAudit()
}

func (b *bill) Change() (*pogo.Bill, *pogo.Bill) {
	return b.before, b.Bill
}

func (b *bill) SetPrice(price int)  {
	b.Price = price
}

func (b *bill) GetPtr() *pogo.Bill {
	return b.Bill
}

func (b *bill) StartAudit() {
	if b.ID == 0 {
		b.before = nil
	} else {
		b.before = &pogo.Bill{}
		copier.Copy(b.before, b)
	}
}
