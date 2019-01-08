package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
	"github.com/jinzhu/gorm"
)

var SRV *http.Server

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
	// пропущена авторизация и csrf
	r := mux.NewRouter()
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Start DataConnector\n")
	})
	r.HandleFunc("/ondulin", Ondulin)
	r.HandleFunc("/quit", QuitHandler)

	//http.ListenAndServe(":2603", r)
	SRV = &http.Server{Addr: ":2603", Handler: r}
	if err := SRV.ListenAndServe(); err != http.ErrServerClosed {
		// Error starting or closing listener:
		log.Printf("HTTP server ListenAndServe: %v", err)
	}
	fmt.Println("starting server at :2603")
}
