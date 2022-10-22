package main

import (
	"fmt"
	"log"
	"os"
	"time"
)

var data = []int64{} // data
var moves = []bool{} // moves
var n int            // elements qty
var l, r int         // heads positions
var buf = []int64{}  // subsequense
var max int64        // local maximum
var answer string    // string to print

func reader() {
	fmt.Scan(&n)
	fmt.Scan(&n)
	// read slice
	data = make([]int64, n)
	for i := 0; i < n; i++ {
		fmt.Fscan(os.Stdin, &data[i])
	}
	fmt.Scan(&n)
	moves = make([]bool, n)
	for i := 0; i < n; i++ {
		var a string
		fmt.Fscan(os.Stdin, &a)
		if a == "R" {
			moves[i] = true
		} else {
			moves[i] = false
		}
	}
}
func findmax() {
	max = -1000000001
	for _, val := range buf {
		if val > max {
			max = val
		}
	}
}
func printtxt() {
	if err := os.WriteFile("files/output.txt", []byte(answer), 0666); err != nil {
		log.Fatal(err)
	}
}
func main() {
	start := time.Now() // count working time
	// read incomes
	reader()
	duration := time.Since(start) // count working time
	// edge case
	l = 0
	r = 0
	max = data[0]
	buf = make([]int64, 1)
	buf[0] = data[0]
	// do
	for _, move := range moves {
		if move {
			r++
			buf = append(buf, data[r])
			if data[r] > max {
				max = data[r]
			}
		} else {
			l++
			if buf[0] == max {
				buf = buf[1:]
				findmax()
			} else {
				buf = buf[1:]
			}
		}
		answer += fmt.Sprintf(" %d", max)
	}
	duration2 := time.Since(start) // count working time
	//fmt.Println(answer)
	printtxt()
	fmt.Println(duration, duration2)
}
