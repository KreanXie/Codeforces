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
		var x int
		fmt.Fscan(in, &x)
		ret := [32]int{}

		r := 0
		for bit := 0; bit < 32; bit++ {
			if x&(1<<bit) != 0 {
				ret[bit] = 1
				r = bit
			}
		}
		// fmt.Println(ret)
		cnt := 0
		l := -1
		for i := 0; i < 32; i++ {
			if ret[i] == 1 {
				cnt++
			} else {
				if cnt > 1 {
					ret[i] = 1
					r = max(r, i)
					for j := l + 2; j < i; j++ {
						ret[j] = 0
					}
					ret[l+1] = -1
					i--
				}
				cnt = 0
				l = i
			}
		}
		fmt.Fprintln(out, r+1)
		for i := range r + 1 {
			fmt.Fprintf(out, "%d ", ret[i])
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
