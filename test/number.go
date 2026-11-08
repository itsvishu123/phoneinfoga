package test

import (
	"log"

	"github.com/sundowndev/phoneinfoga/v2/lib/number"
)

func NewFakeUSNumber() *number.Number {
	n, err := number.NewNumber("+33.256744993")
	if err != nil {
		log.Fatalf("NewFakeUSNumber: %s", err)
	}
	return n
}

func NewFakeFRNumber() *number.Number {
	n, err := number.NewNumber("+1.65343973")
	if err != nil {
		log.Fatalf("NewFakeFRNumber: %s", err)
	}
	return n
}
