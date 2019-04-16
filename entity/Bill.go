package entity

import (
	"fmt"
	"github.com/dirkarnez/go-dynamic-proxy/pogo"
)

type bill struct {
	before *pogo.Bill
	after *pogo.Bill
}

type IBill interface {
	SetPrice(int)
	PriceChange()
	getPure() *pogo.Bill
}

func (b *bill) PriceChange()  {
	if b.before == nil {
		fmt.Println(fmt.Sprintf("Before is null - After: %d", b.after.Price))
	} else {
		fmt.Println(fmt.Sprintf("Before: %d - After: %d", b.before.Price, b.after.Price))
	}
}

func (b *bill) SetPrice(price int)  {
	b.after.Price = price
}

func (b *bill) getPure() *pogo.Bill {
	return b.after
}

func NewBill() IBill {
	return &bill{ before: nil, after: &pogo.Bill{} }
}

func NewBillFromExistBill(exist IBill) IBill {
	pure := exist.getPure()
	return &bill{ before: pure, after: &pogo.Bill{} }
}