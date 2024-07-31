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
		var n int
		fmt.Fscan(in, &n)
		var s string
		fmt.Fscan(in, &s)
		ret := 0
		for i := range s {
			if s[i] == '.' {
				continue
			} else if s[i] == '@' {
				ret++
			} else if s[i] == '*' {
				if i < n-1 && s[i+1] == '*' {
					break
				}
			} else {
				panic("Unknown char")
			}
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
