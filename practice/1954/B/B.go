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
		var n int
		fmt.Fscan(in, &n)
		nums := make([]int, n)
		flag := true
		for i := range nums {
			fmt.Fscan(in, &nums[i])
			if nums[i] != nums[0] {
				flag = false
			}
		}
		num0 := nums[0]
		l, r := 0, 0
		minl := sup
		if flag {
			fmt.Fprintln(out, -1)
			goto next
		}
		if nums[0] != nums[n-1] {
			fmt.Fprintln(out, 0)
			goto next
		}
		for i := 0; i < n-1; i++ {
			if nums[i] != num0 && nums[i+1] != num0 {
				fmt.Fprintln(out, 0)
				goto next
			}
		}
		for r < n {
			for r < n && nums[r] == num0 {
				r++
			}
			minl = min(minl, r-l)
			r++
			l = r
		}
		fmt.Fprintln(out, minl)
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
