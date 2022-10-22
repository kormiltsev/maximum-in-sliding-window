package main

import (
	"fmt"
	"math/rand"
	"time"
)

var n = 100000     // from 1 to 100000
var st = []int64{} // from -1000000000 to 1000000000
var m = 150000     // from 1 to 2n-2
var lr string      // R or L only, delimer = space

var size int64 = 2000000000
var r, l int

func randomer() {
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < n; i++ {
		st = append(st, rand.Int63n(size)-1000000000)
	}
	i := 0
	lr = "R"
	for i < m {
		if rand.Intn(2) == 1 {
			if r < n {
				lr += " R"
				r++
				i++
			}
		} else {
			if l < r {
				lr += " L"
				l++
				i++
			}
		}
	}
}

//	func RandBool() bool {
//		rand.Seed(time.Now().UnixNano())
//		return rand.Intn(2) == 1
//	}
func main() {
	randomer() //v2
	// 1 line: number of elements
	fmt.Println(n)
	// 2 line: list of values
	fmt.Println(st)
	// 3 line: quantity of moves
	fmt.Println(m)
	// 4 line: list of moves (L&R)
	fmt.Println(lr)
}
