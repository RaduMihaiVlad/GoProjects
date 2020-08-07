package main

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

type TestCardString struct {
	inputCard Card
	expected  		string
}

var testsString = []TestCardString {
	{Card {Suit:Diamonds, Value:Two}, "Two of Diamonds"},
	{Card {Suit:Clubs, Value:Ace}, "Ace of Clubs"},
	{Card {Suit:Hearts, Value:Queen}, "Queen of Hearts"},
}
func TestStringCard(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {
		convey.Convey("We should obtain the same value as expected value", func() {
			for _, test := range(testsString) {	
				convey.So(test.inputCard.String(), convey.ShouldEqual, test.expected)
			}
		})
	})
}
