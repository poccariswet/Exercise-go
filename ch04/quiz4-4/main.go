package main

import "fmt"

func rotate(s []int, count int) {
	var a []int = make([]int, count, count*2) // 長さcount=3, キャパcap=6
	copy(a[0:], s[:count])                    // a[0], a[1], a[2] = s[0], s[1], s[2]をコピー
	copy(s[0:], s[count:])                    // s[0], s[1], s[2] = s[3], s[4], s[5]をコピー
	copy(s[len(s)-len(a):], a[:count])        // s[3], s[4], s[5] = a[0], a[1], a[2]をコピー
}

func main() {
	list := []int{0, 1, 2, 3, 4, 5}
	rotate(list, 3)
	fmt.Println(list)
}
