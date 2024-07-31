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
		var a, b, c, d int
		fmt.Fscan(in, &a, &b, &c, &d)
		nums := []int{a, b, c, d}
		sort.Ints(nums)
		if (min(a, b) == nums[0] && max(a, b) == nums[1]) || (min(a, b) == nums[2] && max(a, b) == nums[3]) || (min(a, b) == nums[0] && max(a, b) == nums[3]) || (min(a, b) == nums[1] && max(a, b) == nums[2]) {
			fmt.Fprintln(out, "NO")
		} else {
			fmt.Fprintln(out, "YES")
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
