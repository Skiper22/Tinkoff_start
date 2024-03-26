package main

import (
	"fmt"
)

const N = 300010

var (
	n              int
	k              int
	cost           [N]int
	nameComp       [N]int
	edges          [N][]int
	subtreeMask    [N]int64
	ans            [N]int64
)

func diff(v, p int) {
	subtreeMask[v] = 1 << uint(nameComp[v])
	ans[v] += int64(cost[v])
	for _, to := range edges[v] {
		if to == p {
			continue
		}
		diff(to, v)
		subtreeMask[v] |= subtreeMask[to]
		ans[v] += ans[to]
	}
}

func main() {
	fmt.Scan(&n, &k)
	idsComp := make(map[string]int)
	for i := 0; i < k; i++ {
		s := ""
		fmt.Scan(&s)
		idsComp[s] = i
	}
	root := -1
	for i := 0; i < n; i++ {
		p := 0
		s := ""
		fmt.Scan(&p, &cost[i], &s)
		nameComp[i] = idsComp[s]
		if p == 0 {
			root = i
		} else {
			edges[p-1] = append(edges[p-1], i)
			edges[i] = append(edges[i], p-1)
		}
	}
	diff(root, root)
	var res int64 = -1
	for v := 0; v < n; v++ {
		if subtreeMask[v] == (1<<uint(k))-1 {
			if res > ans[v] || res == -1 {
				res = ans[v]
			}
		}
	}
	fmt.Println(res)
}
