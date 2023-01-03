package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

func main() {
	sc.Split(bufio.ScanWords)
	N, Q := scanInt(), scanInt()

	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	L := make([]int, Q)
	R := make([]int, Q)
	for i := 0; i < Q; i++ {
		L[i], R[i] = scanInt(), scanInt()
	}

	S := make([]int, N+1)
	S[0] = 0
	for i := 1; i < N+1; i++ {
		S[i] = S[i-1] + A[i-1]
	}

	for i := 0; i < Q; i++ {
		fmt.Println(S[R[i]] - S[L[i]-1])
	}
}
