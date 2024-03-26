package main
import (
	"fmt"
)

func maxValue(a int, b int) int{
	if (a >= b){
		return a
	}
	return b
}

func main() {
	var n,m int
	fmt.Scan(&n, &m)
	arr := make([]int, n)
	i := 0
	for i = 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	mx := m
	for i = 0; i < n; i++{
		lMx := arr[i] - 1
		rMx := mx - arr[i]
		mx = maxValue(lMx, rMx)
	}
	fmt.Print(mx)
}
