package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

type Category struct {
	ID   int    `xml:"id,attr" gorm:"primary_key"`
	Name string `xml:",chardata"`
}
type Vendor struct {
	Name string `xml:"vendor"`
}

type Goods struct {
	//gorm.Model
	ID        int     `sql:"AUTO_INCREMENT" gorm:"primary_key"`
	Scu       string  `xml:"id,attr"`
	Available bool    `xml:"available,attr" gorm:"-"`
	Name      string  `xml:"name"`
	Price     float32 `xml:"price" gorm:"-"`
	CatID     int     `xml:"categoryId"`
	Desc      string  `xml:"description"`
	Image     string  `xml:"picture"`
	Vendor    string  `xml:"vendor"`
	Param     []struct {
		Text string `xml:",chardata"`
		Name string `xml:"name,attr"`
		Unit string `xml:"unit,attr"`
	} `xml:"param" gorm:"-"`
}
type Data struct {
	gorm.Model `gorm:"-"`
	ProductId  int `gorm:"Column product_id; primary_key"`
	Quantity   float32
	Price      float32
	Available  bool
	DateAdd    time.Time `gorm:"Column date_add"`
}

func main() {
	// в целям упрощения примера пропущена авторизация и csrf
	r := mux.NewRouter()
	r.HandleFunc("/", Ondulin)
	/*
		r.HandleFunc("/items", handlers.List).Methods("GET")
		r.HandleFunc("/items/new", handlers.AddForm).Methods("GET")
		r.HandleFunc("/items/new", handlers.Add).Methods("POST")
		r.HandleFunc("/items/{id}", handlers.Edit).Methods("GET")
		r.HandleFunc("/items/{id}", handlers.Update).Methods("POST")
		r.HandleFunc("/items/{id}", handlers.Delete).Methods("DELETE")
	*/
	fmt.Println("starting server at :2603")
	http.ListenAndServe(":2603", r)
}

func Ondulin() {
	fmt.Println("Начал загрузку")
	url := "http://store.onduline.com.ua/custom_export.php"
	resp, err := http.Get(url)
	if err != nil {
		fmt.Println("error happend", err)
		return
	}
	defer resp.Body.Close()

	respBody, err := ioutil.ReadAll(resp.Body)
	ParsingVendor(respBody)
	go ParsingCategory(respBody)
	ParsingGoods(respBody, true)
}
