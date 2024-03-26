package main
import (
	"fmt"
	//"strings"
)

func isCorrect(s string) int{
	Tcount := 0
	Icount := 0
	Ncount := 0
	Kcount := 0
	Ocount := 0
	Fcount := 0
	if len(s) != 7{
		return 0
	}
	for _,value := range s{
		switch value {
		case 'T':
			Tcount++
		case 'I':
			Icount++
		case 'N':
			Ncount++
		case 'K':
			Kcount++
		case 'O':
			Ocount++
		case 'F':
			Fcount++
		}
	if (Tcount*Icount*Ncount*Kcount*Ocount == 1) && (Fcount == 2){
		return 2
	}
	}
	return 0
}

func main() {
	var word [100]string
	var size int
	fmt.Scan(&size)
	i := 0
	for i = 0; i < size; i++{
	 	fmt.Scan(&word[i])
	}
	for i = 0; i < size; i++{
		if isCorrect(word[i]) == 2{
			fmt.Printf("Yes\n")
		} else {
			fmt.Printf("No\n")
		}
	}
}
