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
		nums := make([]int, n)
		sum := 0
		for i := range nums {
			fmt.Fscan(in, &nums[i])
			sum += nums[i]
		}
		aver := sum / n
		last := nums[0] - aver
		if nums[0] < aver {
			fmt.Fprintln(out, "NO")
			goto next
		}
		for i := 1; i < n; i++ {
			nums[i] += last
			if nums[i] < aver {
				fmt.Fprintln(out, "NO")
				goto next
			}
			last = nums[i] - aver
		}
		fmt.Fprintln(out, "YES")
		goto next
	next:
		ex--
	}
	out.Flush()
}

func main() {
	solve(os.Stdin, os.Stdout)
}
