package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

const (
	sup = 1 << 30
	inf = -sup
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)
	type pair struct {
		a, b int
	}
	var ex int
	fmt.Fscan(in, &ex)
	for ex != 0 {
		var n, k, q int
		fmt.Fscan(in, &n, &k, &q)
		t := make([]pair, k)
		for i := 0; i < k; i++ {
			fmt.Fscan(in, &t[i].a)
		}
		for i := 0; i < k; i++ {
			fmt.Fscan(in, &t[i].b)
		}
		for q != 0 {
			var d int
			fmt.Fscan(in, &d)
			idx := sort.Search(k, func(i int) bool {
				return t[i].a >= d
			})
			if idx == 0 {
				fmt.Fprintf(out, "%d ", d*t[idx].b/t[idx].a)
			} else {
				fmt.Fprintf(out, "%d ", t[idx-1].b+(d-t[idx-1].a)*(t[idx].b-t[idx-1].b)/(t[idx].a-t[idx-1].a))
			}
			q--
		}
		fmt.Fprintln(out)
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
