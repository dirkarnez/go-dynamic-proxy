package main

import (
	"github.com/dirkarnez/go-dynamic-proxy/entity"
)



func main() {
	bill1 := entity.Factory.NewBill()
	bill1.SetPrice(100)
	bill1.SetPrice(200)
	bill1.Change()
	//
	//bill2 := entity.Factory.NewBillFromExistBill(bill1)
	//bill2.SetPrice(300)
	//bill2.SetPrice(400)
	//bill2.Change()
}