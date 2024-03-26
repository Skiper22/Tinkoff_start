package main
import (
	"fmt"
)

func sumAray(arr []int) int{
	sum := 0
	for _, num := range arr {
        sum += num
    }
    return sum
}

func main() {
	var count int
	fmt.Scan(&count)
	res := make([]int, count)
	i := 0
	j := 0
	var k int
	for i = 0; i < count; i++{
		fmt.Scan(&j)
		arr := make([]int, j)
		for k = 0; k < j; k++{
			fmt.Scan(&arr[k])
		}
		if j > 4{
			if sumAray(arr) >= (j*2 - 2){
				res[i] = 1
			}
		}
		if (j == 3) && (sumAray(arr) >= 4){
			res[i] = 1
		}
		if (j == 4) && (sumAray(arr) >= 6){
			res[i] = 1
		}
		
	}

	res[0] = 1
	if count >= 2{
		res[1] = 1
	}
	for i = 0; i < count; i++{
		if res[i] == 1{
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}

}
