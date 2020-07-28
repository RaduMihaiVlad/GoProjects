package main 

import (
	"golang.org/x/tour/tree"
	"testing"
	"github.com/smartystreets/goconvey/convey"
)

func TestSame(t *testing.T) {
	if Same(tree.New(2), tree.New(2)) != true {
		t.Error("Wrong answer")
	}
}

type TestElement struct {
	x 		 int
	y		 int
	expected bool
}

var tests = []TestElement {
	{1, 2, false},
	{2, 3, false},
	{1, 4, false},
	{2, 2, true},
	{3, 3, true},
	{100, 100, true},
	{15, 2, false},
}

func TestSameArr(t *testing.T) {
	for _, elem := range(tests) {
		if Same(tree.New(elem.x), tree.New(elem.y)) != elem.expected {
			t.Error("Wrong")
		}
	}
}

func TestConvey(t *testing.T) {
	convey.Convey("For all the given tests from array", t, func() {
			
			convey.Convey("We should obtain the same value as expected value", func() {
				for _, test := range(tests) {
					convey.So(Same(tree.New(test.x), tree.New(test.y)), convey.ShouldEqual, test.expected)
				}
		})
	})
}