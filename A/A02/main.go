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

	N, X := scanInt(), scanInt()

	for i := 0; i < N; i++ {
		if scanInt() == X {
			fmt.Println("Yes")
			return
		}
	}
	fmt.Println("No")
}
