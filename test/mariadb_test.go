package test

import (
	"fmt"
	"github.com/dirkarnez/go-dynamic-proxy/entity"
	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	"testing"
)

func TestMariaDB(t *testing.T) {
	var dbDSN = "root:password@tcp(localhost:3306)/gorm?charset=utf8&parseTime=True"
	db, err := gorm.Open("mysql", dbDSN)
	defer db.Close()


	if err != nil {
		t.Error(err)
		return
	}

	// db.Callback().Create()

	bill1 := entity.Factory.NewBill()
	bill1.StartAudit()
	bill1.SetPrice(999)
	fmt.Println(db.Create(bill1.GetPtr()).GetErrors())
	_, newB1 := bill1.Change()
	fmt.Println(newB1.Price)
	bill2 := entity.Factory.NewBill()
	db.First(bill2.GetPtr(), 32)
	bill2.StartAudit()
	bill2.SetPrice(207)
	db.Save(bill2.GetPtr())
	oldB2, newB2 := bill2.Change()
	fmt.Println(oldB2.Price, newB2.Price)
}