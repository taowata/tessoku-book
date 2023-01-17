package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var sc = bufio.NewScanner(os.Stdin)

func scanString() string {
	sc.Scan()
	return sc.Text()
}

func scanInt() int {
	sc.Scan()
	i, err := strconv.Atoi(sc.Text())
	if err != nil {
		panic(err)
	}
	return i
}

type Stack struct {
	st  []string
	top int
}

func NewStack() *Stack {
	return &Stack{
		st:  make([]string, 0),
		top: -1,
	}
}

func (s *Stack) push(x string) {
	s.st = append(s.st, x)
	s.top++
}
func (s *Stack) isEmpty() bool {
	return len(s.st) == 0
}

func (s *Stack) pop() {
	if s.isEmpty() {
		panic("stack is empty.")
	}
	s.st = s.st[:s.top]
	s.top--
}

func (s *Stack) showTop() string {
	if s.isEmpty() {
		panic("stack is empty.")
	}
	return s.st[s.top]
}

func main() {
	sc.Split(bufio.ScanWords)
	Q := scanInt()

	st := NewStack()
	for i := 1; i <= Q; i++ {
		q := scanString()
		switch q {
		case "1":
			st.push(scanString())
		case "2":
			fmt.Println(st.showTop())
		case "3":
			st.pop()
		}
	}
}
