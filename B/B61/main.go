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

func maxIn(s []int) (i int, v int) {
	i = 0
	v = s[0]
	for j := 1; j < len(s); j++ {
		if s[j] > v {
			i = j
			v = s[j]
		}
	}
	return i, v
}

func main() {
	sc.Split(bufio.ScanWords)
	N, M := scanInt(), scanInt()
	A := make([]int, M+1)
	B := make([]int, M+1)
	G := make([]int, N+1)
	for i := 1; i <= M; i++ {
		A[i] = scanInt()
		B[i] = scanInt()
		G[A[i]]++
		G[B[i]]++
	}

	i, _ := maxIn(G)
	fmt.Println(i)
}
