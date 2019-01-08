package main

import (
	"context"
	"fmt"
	"io"
	"io/ioutil"
	"log"
	"net/http"
)

func Ondulin(w http.ResponseWriter, r *http.Request) {
	io.WriteString(w, "Начал загрузку Ондулин\n")

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
	io.WriteString(w, "Закончил загрузку Ондулин\n")
}

func QuitHandler(w http.ResponseWriter, r *http.Request) {
	log.Printf("main: stopping HTTP server")
	io.WriteString(w, "Server shutdown\n")
	// timeout could be given instead of nil as a https://golang.org/pkg/context/
	if err := SRV.Shutdown(context.Background()); err != nil {
		panic(err) // failure/timeout shutting down the server gracefully
	}

}
