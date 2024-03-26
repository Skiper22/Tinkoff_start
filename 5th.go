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
	out := make([]int, 1000)
	l := 0
	var n,m,q int
	fmt.Scan(&n,&m,&q)
	countMan := make([]int, n)
	var i int
	for i = 0; i < n; i++{
		fmt.Scan(&countMan[i])
	}
	friendsA := make([]int, m)
	friendsB := make([]int, m)
	for i = 0; i < m; i++{
		fmt.Scan(&friendsA[i], &friendsB[i])
	}
	var str string
	var v,x int
	for i = 0; i < q; i++{
		// fmt.Println(i)
		fmt.Scan(&str)
		// time.Sleep(1 * time.Second)
		if str == "+"{
			fmt.Scan(&v, &x)
			arr, len := whoFriends(friendsA, friendsB, v)
			// fmt.Print(arr)
			r := 0
			for r = 0; r < len; r++{
			 	countMan[arr[r]-1] += x
			}
		}
		if str == "?"{
			fmt.Scan(&v)
			out[l] = countMan[v-1]
			l++
		}
	}
	for i = 0; i < l; i++{
		fmt.Println(out[i])
	}
}

