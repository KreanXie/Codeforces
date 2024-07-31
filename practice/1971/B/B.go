package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"
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
		if strings.Count(s, s[:1]) == len(s) {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
			ss := []byte(s)
			for i := range ss {
				for j := range ss {
					if i != j && ss[i] != ss[j] {
						ss[i], ss[j] = ss[j], ss[i]
						goto out
					}
				}
			}
		out:
			fmt.Fprintln(out, string(ss))
		}
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
