// 一ヶ月未満、一年未満、一年以上の期間で分類された結果を報告するように修正
package main

import (
	"fmt"
	"log"
	"os"

	"gopl.io/ch4/github"
)

var 

func SortCreatedAt() {

}

func main() {
	result, err := github.SearchIssues(os.Args[1:])
	if err != nil {
		log.Fatal(err)
	}
	fmt.Printf("%d issues:\n", result.TotalCount)
	for _, item := range result.Items {
		fmt.Printf("#%-5d %9.9s %.55s\n",
			item.Number, item.User.Login, item.Title)
		fmt.Println(item.CreatedAt)
		fmt.Println()
	}
}



// fmt.Sprintf("#%-5d %9.9s %.55s\n", item.Number, item.User.Login, item.Title)
