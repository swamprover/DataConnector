package main

import "testing"

func TestParsingGoods(t *testing.T) {
	type args struct {
		dataXML []byte
		toBase  bool
	}
	tests := []struct {
		name string
		args args
	}{

		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParsingGoods(tt.args.dataXML, true)
		})
	}
}
