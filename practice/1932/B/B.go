package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

func solve(_r io.Reader, _w io.Writer) {
	in := bufio.NewReader(_r)
	out := bufio.NewWriter(_w)

	var ex int64
	fmt.Fscan(in, &ex)
	for ex != 0 {
		var n int
		fmt.Fscan(in, &n)
		a := make([]int, n)
		for i := range a {
			fmt.Fscan(in, &a[i])
		}
		cur := a[0]
		for i := 1; i < n; i++ {
			// fmt.Fprintln(out, ex, cur)
			if cur >= a[i] {
				a[i] *= sort.Search(1e9, func(m int) bool {
					return m*a[i] > cur
				})
			}
			cur = a[i]
		}
		fmt.Fprintln(out, cur)
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
