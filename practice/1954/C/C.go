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
		var x, y []byte
		fmt.Fscan(in, &x, &y)
		n := len(x)
		flag := 0
		idx := 0
		for i := 0; i < n; i++ {
			idx = i
			if x[i] > y[i] {
				flag = 1
				break
			} else if x[i] < y[i] {
				flag = 2
				break
			}
		}
		switch flag {
		case 0:
			break
		case 1:
			for i := idx + 1; i < n; i++ {
				if y[i] < x[i] {
					y[i], x[i] = x[i], y[i]
				}
			}
		case 2:
			for i := idx + 1; i < n; i++ {
				if x[i] < y[i] {
					x[i], y[i] = y[i], x[i]
				}
			}
		default:
			panic("unknown situation")
		}
		fmt.Fprintln(out, string(x))
		fmt.Fprintln(out, string(y))
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
