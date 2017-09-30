//URLを平行に取り出して、時間と大きさを表示
/*
大量のデータを生成するウェブサイトを見つけなさい。報告される時間が大きく変化するか調べるために'fetchall'を2回続けて実行して、キャッシュされているかどうか調査しなさい。
毎回同じ内容を得ているでしょうか。'fetchall'を修正して、その出力をファイルへ保存するようにして調べられるようにしなさい
*/
package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"time"
)

func fetch(url string, ch chan<- string) { //送信用チャネル
	start := time.Now()
	resp, err := http.Get(url)
	if err != nil {
		ch <- fmt.Sprint(err) //chチャネルへ送信
	}
	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close()
	if err != nil {
		ch <- fmt.Sprintf("while reading %s: %v\n", url, err)
		return
	}
	secs := time.Since(start).Seconds()
	ch <- fmt.Sprintf("%.2fs %7d %s\n", secs, nbytes, url)
}

func main() {
	start := time.Now()
	ch := make(chan string)
	file, err := os.Create("out.txt")
	defer file.Close()
	if err != nil {
		fmt.Printf("file create error: %v\n", err)
		return
	}
	for _, url := range os.Args[1:] {
		go fetch(url, ch) // ゴルーチンを開始
	}
	for range os.Args[1:] {
		//fmt.Println(<-ch) //chチャネルから受信
		file.Write(([]byte)(<-ch))
	}
	//fmt.Printf("%.2fs elapsed\n", time.Since(start).Seconds())
	output := fmt.Sprintf("%.2fs elapsed\n", time.Since(start).Seconds())
	file.Write(([]byte)(output))
}
