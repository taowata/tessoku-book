package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	r := bufio.NewReader(os.Stdin)
	w := bufio.NewWriter(os.Stdout)
	defer w.Flush()

	var a, b int

	fmt.Fscan(r, &a, &b)

	fmt.Fprintln(w, a+b)
}
