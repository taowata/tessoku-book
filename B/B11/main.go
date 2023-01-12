package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
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

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}
	Q := scanInt()
	X := make([]int, Q)
	for i := 0; i < Q; i++ {
		X[i] = scanInt()
	}

	sort.Ints(A)

	for j := 0; j < Q; j++ {
		xi := sort.Search(len(A), func(i int) bool { return A[i] >= X[j] })
		fmt.Println(xi)
	}
}
