package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
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

		goto next
	next:
		ex--
	}
	rw.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
