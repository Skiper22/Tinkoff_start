package main

import (
	"fmt"
	// "strconv"
	// "strings"
	// "time"
)

// func extractNumbers(str string) []int {
//   // str := "+ 1 3 5 6 8 42"
//   strs := strings.Split(str, " ")
//   var ints []int
//   for _, s := range strs {
//     num, err := strconv.Atoi(s)
//     if err == nil {
//       ints = append(ints, num)
//     }
//   }

//   return ints
// }

func whoFriends(arr1 []int, arr2 []int, num int) ([]int, int){
	res := make([]int, 1000)
	k := 0
	for i, value := range arr1{
		if value == num{
			res[k] = arr2[i]
			k++
		}
		if arr2[i] == num{
			res[k] = value
			k++
		}
	}
	return res, k
}

func main() {
	//out := make([]int, 1000)
	var n,q int
	fmt.Scan(&n,&q)
	arr := make([]int, n)
	var i int
	for i = 0; i < n; i++{
		fmt.Scan(&arr[i])
	}
	var str string
	var l,r,x int
	for i = 0; i < q; i++{
		// fmt.Println(i)
		fmt.Scan(&str)
		// time.Sleep(1 * time.Second)
		if str == "+"{
			j := 0
			fmt.Scan(&l, &r, &x)
			for j = l; j < r; j++{
			 	arr[j] += x
			}
		}
		var k,b int
		if str == "?"{
			fmt.Scan(&l, &r, &k, &b)

		}
	}
	// for i = 0; i < l; i++{
	// 	fmt.Println(out[i])
}

