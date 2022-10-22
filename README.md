# Maximum in sliding window
Array of "n" values int64.
Two pointers "l" and "r" at Array[0]. One of them moves 1 position right per step. Index of left pointer can be <= right. 

1<=n<=100 000

-1 000 000 000 <= a(i) <= 1 000 000 000


Income (stdin or input.txt):

n //qty of elements

a1 a2 a3 ...  // data

m // qty of moves m <= 2n-2

R L R R L ... // list of moves

Outcome (stdout or output.txt):

A1 A2 A3 ... // Ai = max in subsequense a1 <= l <= a(i) a(i+1) ... <= r <= a(n). Qty = m.

Goal: time <= 0.6 sec  memory <= 64Mb