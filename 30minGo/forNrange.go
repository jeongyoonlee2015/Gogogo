//반복문에서 range문을 쓰면 슬라이스, 맵을 이터레이트할 수 있다. (전부 for로 바꿔보기) 

package main

import "fmt"

var two = []int{1, 2, 4, 8, 16, 32, 64, 128}
var thr = []int{1, 3, 9, 27, 81, 243, 729, 2187}

func main() {
	fmt.Println("<<2의 제곱>>")
	for i, v := range two {
		fmt.Printf("2^%d = %d\n", i, v)
	}
	fmt.Println("<<3의 제곱>>")
	for i, v := range thr {
		fmt.Printf("3^%d = %d\n", i, v)
	}
}
