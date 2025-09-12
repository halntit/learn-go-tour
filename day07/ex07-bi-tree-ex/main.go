package main

import (
	"fmt"

	"golang.org/x/tour/tree"
)

// Walk walks the tree t sending all values
// from the tree to the channel ch.
func Walk(t *tree.Tree, ch chan int) {
	if t == nil {
		return
	}
	// Walk(t.Left, ch)  // traverse left
	ch <- t.Value // send current node
	// Walk(t.Right, ch) // traverse right
}

// WalkWrapper calls Walk and then closes channel
func WalkWrapper(t *tree.Tree, ch chan int) {
	Walk(t, ch)
	close(ch) // important: signal no more values
}

// Same determines whether the trees
// t1 and t2 contain the same values.
func Same(t1, t2 *tree.Tree) bool {
	ch1, ch2 := make(chan int), make(chan int)

	go WalkWrapper(t1, ch1)
	go WalkWrapper(t2, ch2)

	for {
		v1, ok1 := <-ch1
		v2, ok2 := <-ch2
		if ok1 != ok2 { // channels ended at different times
			return false
		}
		if !ok1 { // both closed
			break
		}
		if v1 != v2 {
			return false
		}
	}
	return true
}

func main() {
	// 1. Test Walk
	ch := make(chan int)
	go WalkWrapper(tree.New(1), ch)
	for v := range ch {
		fmt.Println(v) // should print 1..10
	}

	// 2. Test Same
	fmt.Println("Same(tree.New(1), tree.New(1)) =", Same(tree.New(1), tree.New(1))) // true
	fmt.Println("Same(tree.New(1), tree.New(2)) =", Same(tree.New(1), tree.New(2))) // false

}
