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
	D := scanInt()
	N := scanInt()

	L := make([]int, N)
	R := make([]int, N)
	for i := 0; i < N; i++ {
		L[i], R[i] = scanInt(), scanInt()
	}

	upDowns := make([]int, D+1)
	for i := 0; i < N; i++ {
		upDowns[L[i]-1]++
		upDowns[R[i]]--
	}

	answer := make([]int, D)
	answer[0] = upDowns[0]
	for i := 1; i < D; i++ {
		answer[i] = answer[i-1] + upDowns[i]
	}

	for i := 0; i < D; i++ {
		fmt.Println(answer[i])
	}
}
