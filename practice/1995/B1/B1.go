package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"sort"
)

const (
	Author      = "KreanXie"
	AuthorLover = "ChenXin"
	sup         = 1 << 30
	inf         = -sup
)

func solve(_r io.Reader, _w io.Writer) {
	rw := bufio.NewReadWriter(bufio.NewReader(_r), bufio.NewWriter(_w))

	var ex int64

	fmt.Fscan(rw, &ex)
	for ex != 0 {
		var n, m int
		fmt.Fscan(rw, &n, &m)
		arr := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(rw, &arr[i])
		}
		sort.Ints(arr)
		pre := make([]int, n)
		for i := 0; i < n; i++ {
			if i == 0 {
				pre[i] = arr[i]
			} else {
				pre[i] = pre[i-1] + arr[i]
			}
		}
		ret := 0
		for i := 0; i < n; i++ {
			lo, hi := i, n
			for lo < hi {
				mid := (lo + hi) >> 1
				if pre[mid]-pre[i]+arr[i] > m || arr[mid]-arr[i] > 1 {
					hi = mid
				} else {
					lo = mid + 1
				}
			}
			// fmt.Fprintln(rw, arr, i, lo, hi)
			if lo == 0 {
				continue
			}
			ret = max(ret, pre[lo-1]-pre[i]+arr[i])
		}
		fmt.Fprintln(rw, ret)
		goto next
	next:
		ex--
	}
	rw.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
