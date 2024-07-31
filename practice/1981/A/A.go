package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

const (
	sup = 1 << 30
	inf = -sup
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	yes := func() {
		fmt.Fprintln(out, "YES")
	}
	no := func() {
		fmt.Fprintln(out, "NO")
	}
	_, _ = yes, no
	var ex int64
	fmt.Fscan(in, &ex)
	for ex != 0 {
		var l, r int
		fmt.Fscan(in, &l, &r)
		bit := 1
		ret := 0
		for bit < 32 {
			if (1<<bit)&r != 0 {
				ret = bit
			}
			bit++
		}
		fmt.Fprintln(out, ret)
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
