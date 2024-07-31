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
		var n int
		fmt.Fscan(in, &n)
		nums := make([]int, n)
		for i := 0; i < n; i++ {
			fmt.Fscan(in, &nums[i])
		}
		ret := 0
		for i := 0; i < n-1; i++ {
			ret = max(ret, min(nums[i], nums[i+1]))
		}
		for i := 0; i < n-2; i++ {
			arr := []int{nums[i], nums[i+1], nums[i+2]}
			sort.Ints(arr)
			ret = max(ret, arr[1])
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
