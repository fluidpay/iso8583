package main

import (
	"fmt"
	"github.com/fluidpay/iso8583"
	"os"
	"reflect"
	"strings"
)

var actual iso8583.Message

type Printer interface {
	String() string
}

func main() {
	for {
		var msg string
		fmt.Print("> ")
		fmt.Scanf("%s", &msg)

		switch msg {
		case "exit", "e":
			fmt.Println("Bye :)")
			os.Exit(0)
		default:
			if strings.Contains(msg, "search") {
				printActual(strings.Replace(msg, "search ", "", -1))
			} else {
				err := actual.Decode([]byte(msg))
				if err != nil {
					fmt.Println(err.Error())
					os.Exit(1)
				}
				printActual("")
			}
		}
	}
}

func printActual(searchFor string) {
	rv := reflect.ValueOf(actual)
	rt := rv.Type()

	for i := 0; i < rt.NumField(); i++ {
		if searchFor != "" {
			if strings.Contains(strings.ToLower(rt.Field(i).Name), searchFor) {
				if reflect.Indirect(rv.Field(i)).IsValid() {
					fmt.Printf("%s -> %s\n", rt.Field(i).Name, reflect.Indirect(rv.Field(i)))
				}
			}
		} else {
			if reflect.Indirect(rv.Field(i)).IsValid() {
				fmt.Printf("%s -> %s\n", rt.Field(i).Name, reflect.Indirect(rv.Field(i)))
			}
		}
	}
}
