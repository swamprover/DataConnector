package main

import (
	"fmt"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jinzhu/gorm"
)

//AddUpdate добавиить, апргрейдить
func AddUpdateGoods(rec Goods, tabl string) {
	db, err := gorm.Open("mysql", "root:jcnjhj;yj26@tcp(localhost:3306)/profit?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer db.Close()
	//db.LogMode(true) //включение лога
	var result Goods

	if err := db.Table(tabl).Where("scu = ?", rec.Scu).Find(&result).Error; gorm.IsRecordNotFoundError(err) {
		db.Table(tabl).Create(&rec)
		var resData Data
		resData.ProductId = rec.ID
		resData.Price = rec.Price
		resData.Available = rec.Available
		resData.DateAdd = time.Now()
		// создаем потому что data связана с goods foreign key
		db.Table("data").Create(&resData)
	} else {
		db.Table(tabl).Where("scu = ?", rec.Scu).Update(&rec)

		var resData Data
		if err := db.Table("data").Where("product_id = ?", result.ID).Find(&resData).Error; gorm.IsRecordNotFoundError(err) {
			resData.ProductId = result.ID
			resData.Price = rec.Price
			resData.Available = rec.Available
			resData.DateAdd = time.Now()
			db.Table("data").Create(&resData)
		} else {
			resData.ProductId = result.ID
			resData.Price = rec.Price
			resData.Available = rec.Available
			resData.DateAdd = time.Now()
			db.Table("data").Where("product_id = ?", resData.ProductId).Update(&resData)
		}
		//fmt.Println("Found")
	}

}
func AddUpdateCategories(rec Category) {
	db, err := gorm.Open("mysql", "root:jcnjhj;yj26@tcp(localhost:3306)/profit?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer db.Close()

	var result Category
	if err := db.Table("categories").Where("id = ?", rec.ID).Find(&result).Error; gorm.IsRecordNotFoundError(err) {
		db.Table("categories").Create(&rec)
	} else {
		db.Table("categories").Where("id = ?", rec.ID).Update(&rec)
	}
}
