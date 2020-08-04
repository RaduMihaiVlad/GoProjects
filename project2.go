package main

import (
	"golang.org/x/tour/tree"
	"fmt"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	
	if (t.Left != nil) {
		Walk(t.Left, ch)
	}
	ch <- t.Value
	if (t.Right != nil) {
		Walk(t.Right, ch)
	}
}

// Same determines whether the trees
// t1 and t2 contain the same values.

func Same(t1, t2 *tree.Tree) bool {
	c1 := make(chan int)
	c2 := make(chan int)

	go Walk(t1, c1)
	go Walk(t2, c2)

	a1Array := make([]int, 0, 10)
	a2Array := make([]int, 0, 10)

	for len(a1Array) != 10 || len(a2Array) != 10 {
		select {
			case x := <- c1:
				a1Array = append(a1Array, x)
			case y := <- c2:
				a2Array = append(a2Array, y)
		}
	}
	for i := 0; i < 10; i++ {
		if a1Array[i] != a2Array[i] {
			return false
		}
	}
	return true
}

/*
	I am implementing the exercise described here: https://go-tour-ro.appspot.com/#68
	I must tell if two trees have the same values, and, in order to do this, I am using 
	goroutines and channels

*/

func main() {

	/*t := tree.New(2)
	ch := make(chan int)

	go Walk(t, ch)

	for i := 0; i < 10; i++ {
		x := <- ch
		fmt.Println(x)
	}*/

	fmt.Println(Same(tree.New(1), tree.New(1)))
	fmt.Println(Same(tree.New(1), tree.New(2)))

}