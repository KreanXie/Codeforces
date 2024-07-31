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
		if n < 28 {
			fmt.Fprintln(out, "aa"+string(byte(n-3+int('a'))))
		} else {
			if n < 54 {
				fmt.Fprintln(out, "a"+string(byte(n-28+int('a')))+"z")
			} else {
				fmt.Fprintln(out, string(byte(n-53+int('a')))+"zz")
			}
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
