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

	N := scanInt()
	A := make([]int, N)
	for i := 0; i < N; i++ {
		A[i] = scanInt()
	}

	Q := scanInt()
	L := make([]int, Q)
	R := make([]int, Q)
	for i := 0; i < Q; i++ {
		L[i], R[i] = scanInt(), scanInt()
	}

	Atari := make([]int, N+1)
	Hazure := make([]int, N+1)
	Atari[0] = 0
	Hazure[0] = 0

	for i := 1; i < N+1; i++ {
		Atari[i] = Atari[i-1]
		if A[i-1] == 1 {
			Atari[i] += 1
		}
		Hazure[i] = Hazure[i-1]
		if A[i-1] == 0 {
			Hazure[i] += 1
		}
	}

	for i := 0; i < Q; i++ {
		NumAtari := Atari[R[i]] - Atari[L[i]-1]
		NumHazure := Hazure[R[i]] - Hazure[L[i]-1]

		switch {
		case NumAtari > NumHazure:
			fmt.Println("win")
		case NumAtari == NumHazure:
			fmt.Println("draw")
		case NumAtari < NumHazure:
			fmt.Println("lose")
		}
	}
}
