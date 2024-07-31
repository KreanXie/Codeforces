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
		var s string
		fmt.Fscan(in, &s)
		ret := 0
		flag := false
		for i := 0; i < len(s)-1; i++ {
			if s[i] == '0' && s[i+1] == '1' {
				flag = true
			}
		}
		i := 0
		for i < len(s) {
			for i < len(s)-1 && s[i] == s[i+1] {
				i++
			}
			i++
			ret++
		}
		if flag {
			ret--
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
