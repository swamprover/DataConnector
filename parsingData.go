package main

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"io"
)

//ParsingCategory парсим только поставщиков
func ParsingCategory(dataXML []byte) {

	input := bytes.NewReader(dataXML)
	decoder := xml.NewDecoder(input)
	var element Category

	for {
		tok, tokenErr := decoder.Token()
		if tokenErr != nil && tokenErr != io.EOF {
			fmt.Println("error happend", tokenErr)
			break
		} else if tokenErr == io.EOF {
			break
		}
		if tok == nil {
			fmt.Println("t is nil break")
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "category" {

				if err := decoder.DecodeElement(&element, &tok); err != nil {
					fmt.Println("error happend", err)
				}

				fmt.Printf("http.Get body %#v\n", element)
				AddUpdateCategories(element)
			}
			if tok.Name.Local == "offer" {
				break
			}
		}
	}
}

//ParsingVendor парсим только поставщиков
func ParsingVendor(dataXML []byte) {

	input := bytes.NewReader(dataXML)
	decoder := xml.NewDecoder(input)
	var element Vendor

	for {
		tok, tokenErr := decoder.Token()
		if tokenErr != nil && tokenErr != io.EOF {
			fmt.Println("error happend", tokenErr)
			break
		} else if tokenErr == io.EOF {
			break
		}
		if tok == nil {
			fmt.Println("t is nil break")
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "offer" {

				if err := decoder.DecodeElement(&element, &tok); err != nil {
					fmt.Println("error happend", err)
				}

				fmt.Printf("http.Get body %#v\n", element)
			}
		}
	}
}

//ParsingGoods парсим товары
func ParsingGoods(dataXML []byte, toBase bool) {

	input := bytes.NewReader(dataXML)
	decoder := xml.NewDecoder(input)
	var element Goods

	for {
		tok, tokenErr := decoder.Token()
		if tokenErr != nil && tokenErr != io.EOF {
			fmt.Println("error happend", tokenErr)
			break
		} else if tokenErr == io.EOF {
			break
		}
		if tok == nil {
			fmt.Println("t is nil break")
		}
		switch tok := tok.(type) {
		case xml.StartElement:
			if tok.Name.Local == "offer" {

				if err := decoder.DecodeElement(&element, &tok); err != nil {
					fmt.Println("error happend", err)
				}

				if toBase {
					AddUpdateGoods(element, "goods")
				}
			}

		}
	}
}
