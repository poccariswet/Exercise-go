package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"strings"
	"time"
)

type locationTimeWriter struct {
	w            io.Writer
	locationName string
}

func (w *locationTimeWriter) Write(data []byte) (int, error) {
	return w.w.Write(data)
}

func main() {
	flag.Parse()
	fmt.Println(flag.Args())
	for _, r := range flag.Args() {
		var name_addr = strings.Split(r, "=")
		var name = name_addr[0]
		var addr = name_addr[1]
		go handleConn(name, addr)
	}
	for {
		time.Sleep(100 * time.Millisecond)
	}
}

func handleConn(name string, address string) {
	conn, err := net.Dial("tcp", address)
	fmt.Printf("Connected [%s] %s\r\n", name, address)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	var out = &locationTimeWriter{w: os.Stdout, locationName: name}
	mustCopy(out, conn)
}

func mustCopy(dst io.Writer, src io.Reader) {
	if _, err := io.Copy(dst, src); err != nil {
		log.Fatal(err)
	}
}
