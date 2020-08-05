package main 

import (
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

type TestPhoneNumberElement struct {
	inputPhone string
	expected   string
}

var testsNormalize = []TestPhoneNumberElement {
	{"123-456-789", "123456789"},
	{"(0723) 564 234", "0723564234"},
	{"(0712)-123-453", "0712123453"},
}

type TestIsDigitElement struct {
	digit 	rune
	expeted bool
}

var testsIsDigit = []TestIsDigitElement {
	{'1', true},
	{'a', false},
	{'5', true},
	{'4', true},
	{'7', true},
	{'s', false},
	{'5', true},
	{'f', false},
	{'v', false},

}

func TestIsDigit(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {
		convey.Convey("We should obtain the same value as expected value", func() {
			for _, test := range(testsIsDigit) {
				convey.So(IsDigit(test.digit), convey.ShouldEqual, test.expeted)
			}
		})
	})
}

func TestNormalize(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {
			
		convey.Convey("We should obtain the same value as expected value", func() {
			for _, test := range(testsNormalize) {
				convey.So(Normalize(test.inputPhone), convey.ShouldEqual, test.expected)
			}
		})
	})
}