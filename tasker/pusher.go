package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n = 10000       // from 1 to 100000
var st = []int64{}  // from -1000000000 to 1000000000
var m = 12000       // from 1 to 2n-2
var lr = []string{} // R or L only

var size int64
var r, l int

func randomer() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		st = append(st, rand.Int63n(size)-1000000000)
	}
	for i := 0; i < m; i++ {
		if rand.Intn(2) == 1 {
			if r < n {
				lr = append(lr, "R")
				r++
			}
		} else {
			if l < r {
				lr = append(lr, "L")
				l++
			}
		}
	}
}

//	func RandBool() bool {
//		rand.Seed(time.Now().UnixNano())
//		return rand.Intn(2) == 1
//	}
func main() {
	size = 2000000000 // -1000000000-+1000000000
	randomer()        //v2
	// 1 line: number of elements
	fmt.Println(n)
	// 2 line: list of values
	fmt.Println(st)
	// 3 line: quantity of moves
	fmt.Println(m)
	// 4 line: list of moves (L&R)
	fmt.Println(lr)
}
