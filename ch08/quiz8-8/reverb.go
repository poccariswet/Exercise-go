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

const (
	timeout = 10 * time.Second
)

func echo(c net.Conn, shout string, delay time.Duration, wg *sync.WaitGroup) {
	defer (*wg).Done()
	fmt.Fprintln(c, "\t", strings.ToUpper(shout))
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", shout)
	time.Sleep(delay)
	fmt.Fprintln(c, "\t", strings.ToLower(shout))
}

func Scanner(c net.Conn, strch chan<- string) {
	defer close(strch)
	input := bufio.NewScanner(c)
	for input.Scan() {
		strch <- input.Text()
	}
}

func handleConn(c net.Conn) {
	var wg sync.WaitGroup
	strch := make(chan string)
	timech := time.After(timeout)

	go Scanner(c, strch)

	for done := false; !done; {
		select {
		case word := <-strch:
			wg.Add(1)
			go echo(c, word, 1*time.Second, &wg)
			timech = time.After(timeout)
		case <-timech:
			done = true
		}
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
			log.Print(err)
			continue
		}
		go handleConn(conn)
	}
}
