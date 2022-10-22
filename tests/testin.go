//go run ./tests/testin.go | go run ./cmd/task201022/main.go | go run ./tests/testout.go

package main

import "fmt"

var n int       // from 1 to 100000
var st []int64  // from -1000000000 to 1000000000
var m int       // from 1 to 2n-2
var lr []string // R or L only, delimer = space

type ans struct {
	a []int64
}

func nonrandomer() {
	st = []int64{-1000000000, -100, 100, 1, 1, 2, 3, 4, 3, 2, 1, 1000000000}
	lr = []string{"R", "R", "R", "R", "L", "L", "L", "L", "R", "L", "R", "L", "R", "L", "R", "R", "L", "L", "R", "R", "L"}
	//answer := "-100 100 100 100 100 100 1 1 2 2 3 3 4 4 4 4 3 2 2 1000000000 1000000000"
	n = len(st)
	m = len(lr)
}

func main() {
	nonrandomer()
	// 1 line: number of elements
	fmt.Println(n)
	// 2 line: list of values
	fmt.Println(st)
	// 3 line: quantity of moves
	fmt.Println(m)
	// 4 line: list of moves (L&R)
	moves := ""
	for _, move := range lr {
		moves = fmt.Sprintf("%s %s", moves, move)
	}
	fmt.Println(moves)
}
