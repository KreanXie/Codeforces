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

	var ex int64
	fmt.Fscan(in, &ex)
	for ex != 0 {
		var h, n int
		fmt.Fscan(in, &h, &n)
		a, c := make([]int, n), make([]int, n)
		initialSum := 0
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &a[i])
			initialSum += a[i]
		}
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &c[i])
		}
		fmt.Fprintln(out, sort.Search(1e12, func(m int) bool {
			sum := initialSum
			for i := 0; i < n; i++ {
				sum += a[i] * (m / c[i])
			}
			return sum >= h
		})+1)
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
