package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
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
	N, M := scanInt(), scanInt()
	A := make([]int, M+1)
	B := make([]int, M+1)
	G := make([][]string, N+1)
	for i := 1; i <= M; i++ {
		A[i] = scanInt()
		B[i] = scanInt()
		G[A[i]] = append(G[A[i]], strconv.Itoa(B[i]))
		G[B[i]] = append(G[B[i]], strconv.Itoa(A[i]))
	}

	for i := 1; i <= N; i++ {
		fmt.Printf("%d: {%s}\n", i, strings.Join(G[i], ", "))
	}

}
