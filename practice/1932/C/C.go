package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)

	var ex int64
	fmt.Fscan(in, &ex)
	for ex != 0 {
		var n, m int
		fmt.Fscan(in, &n, &m)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		var s string
		fmt.Fscan(in, &s)

		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
