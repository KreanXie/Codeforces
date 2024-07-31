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

	var ex int64
	fmt.Fscan(in, &ex)
	for ex != 0 {
		var x, y int
		fmt.Fscan(in, &x, &y)
		fmt.Fprintln(out, min(x, y), max(x, y))
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
