package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var sc = bufio.NewScanner(os.Stdin)

func main() {
	sc.Split(bufio.ScanWords)
	sc.Scan()

	Ns := strings.Split(sc.Text(), "")

	answer := 0
	lenN := len(Ns)
	for i := 0; i < lenN; i++ {
		kurai := 1 << (lenN - 1 - i)
		var keta int
		if Ns[i] == "0" {
			keta = 0
		}
		if Ns[i] == "1" {
			keta = 1
		}
		answer += keta * kurai
	}
	fmt.Println(answer)
}
