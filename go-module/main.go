package main

import (
	"fmt"
	"go-module/dev"

	"github.com/leekchan/accounting"
)

func main() {
	ac := accounting.Accounting{Symbol: "$", Precision: 2}
	fmt.Println(ac.FormatMoney(123456789.213123))
	dev.Println("xinchaocaccau")
}
