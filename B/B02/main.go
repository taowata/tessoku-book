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

	A, B := scanInt(), scanInt()
	for i := A; i < B+1; i++ {
		if 100%i == 0 {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
