// 一ヶ月未満、一年未満、一年以上の期間で分類された結果を報告するように修正
package main

import (
	"fmt"
	"log"
	"os"
	"strconv"
	"strings"
	"time"

	"gopl.io/ch4/github"
)

/*
一ヶ月未満: 0
一年未満: 1
一年以上: 2
*/

var Month []string
var LessThanYear []string
var MoreThanYear []string

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

func SortCreatedAt(sep [3]int) int {
	t := time.Now()
	now := SeparateTime(t)
	if now[0] == sep[0] && now[1] == sep[1] {
		return 0
	} else if now[0] == sep[0] && now[1]-1 == sep[1] && now[2] <= sep[2] {
		return 0
	} else if now[0]-1 == sep[0] && now[1] <= sep[1] && now[2] <= sep[2] {
		return 1
	} else if now[0]-1 == sep[0] && now[1] < sep[1] {
		return 1
	} else if now[0] == sep[0] && now[1] > sep[1] {
		return 1
	} else {
		return 2
	}
}

func main() {
	var create int
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		sep := SeparateTime(item.CreatedAt)
		create = SortCreatedAt(sep)

		if create == 0 {
			Month = append(Month, fmt.Sprintf("#%-5d %9.9s %.55s\n%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt))
		} else if create == 1 {
			LessThanYear = append(LessThanYear, fmt.Sprintf("#%-5d %9.9s %.55s\n%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt))
		} else if create == 2 {
			MoreThanYear = append(MoreThanYear, fmt.Sprintf("#%-5d %9.9s %.55s\n%s\n", item.Number, item.User.Login, item.Title, item.CreatedAt))
		}

	}

	if len(Month) != 0 {
		fmt.Println()
		fmt.Println("一ヶ月未満のissues")
		for _, month := range Month {
			fmt.Print(month)
		}
	}
	if len(LessThanYear) != 0 {
		fmt.Println()
		fmt.Println("一年未満のissues")
		for _, less := range LessThanYear {
			fmt.Print(less)
		}
	}
	if len(MoreThanYear) != 0 {
		fmt.Println()
		fmt.Println("一年以上のissues")
		for _, more := range MoreThanYear {
			fmt.Print(more)
		}
	}
}

// fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
