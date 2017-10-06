package main

import (
	"fmt"
	"reflect"
	"strconv"
	"strings"
	"time"
)

func main() {
	t := time.Now()
	fmt.Println(reflect.TypeOf(t.UTC()))
	fmt.Printf("%s\n", t.UTC())

	var ss [3]int
	sep := strings.SplitN(fmt.Sprintf("%s", t.UTC()), "-", 3)
	sep2 := strings.Split(sep[2], " ")
	sep[2] = sep2[0]
	for i, s := range sep {
		in, _ := strconv.Atoi(s)
		ss[i] = in
	}
}

func SeparateTime(t time.Time) [3]int {
	var sep [3]int

	str := strings.SplitN(fmt.Sprintf("%s", t.UTC()), "-", 3)
	s := strings.Split(str[2], " ")
	str[2] = s[0]
	for i, s := range str {
		in, _ := strconv.Atoi(s)
		sep[i] = in
	}
	return sep
}
