package main

import (
	"bufio"
	"fmt"
	"log"
	"net"
	"strings"
	"sync"
	"time"
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
	(*wg).Done()
}

func handleConn(c net.Conn) {
	input := bufio.NewScanner(c)
	var wg sync.WaitGroup // 活動中の御ルーチンの数
	for input.Scan() {
		wg.Add(1)
		go echo(c, input.Text(), 1*time.Second, &wg)
	}
	wg.Wait()
	c.(*net.TCPConn).CloseWrite()
}

func main() {
	l, err := net.Listen("tcp", "localhost:8000")
	if err != nil {
		log.Fatal(err)
	}
	for {
		conn, err := l.Accept()
		if err != nil {
			log.Print(err) // e.g., connection aborted
			continue
		}
		go handleConn(conn)
	}
}
