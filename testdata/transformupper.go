package main

import "fmt"

type CamelCaseValue int

const (
	CamelCaseValueOne CamelCaseValue = iota
	CamelCaseValueTwo
	CamelCaseValueThree
)

func main() {
	ck(CamelCaseValueOne, "CAMELCASEVALUEONE")
	ck(CamelCaseValueTwo, "CAMELCASEVALUETWO")
	ck(CamelCaseValueThree, "CAMELCASEVALUETHREE")
	ck(-127, "CAMELCASEVALUE(-127)")
	ck(127, "CAMELCASEVALUE(127)")
}

func ck(value CamelCaseValue, str string) {
	if fmt.Sprint(value) != str {
		panic("transformlowercamel.go: " + str)
	}
}
