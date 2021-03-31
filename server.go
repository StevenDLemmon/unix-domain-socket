package main

import (
	"fmt"
	"log"
	"net"
	"os"
	"strings"
)

const SockAddr = "/tmp/memory.sock"

func server(c net.Conn) {
	// Get the data send by client
	buf := make([]byte, 1024)
	n, err := c.Read(buf[:])
	if err != nil {
		return
	}

	// The client sends over the data in the format of "PID MEMORY\n" so split the incoming data
	// by each PID MEMORY pair
	temp := strings.Split(string(buf[0:n]), "\n")

	for i := 0; i < len(temp) - 1; i++ {
		// Format each PID MEMORY pair into JSON and print to STDOUT
		temp1 := strings.Split(temp[i], " ")
		fmt.Printf("{\"PID\":%s,\"MEMORY\":%s}\n", temp1[0], temp1[1])
	}

	c.Close()
}

func main() {
	// Remove any old sockets that could still be around
	if err := os.RemoveAll(SockAddr); err != nil {
		log.Fatal(err)
	}

	// Start listening to the UDS 
	l, err := net.Listen("unix", SockAddr)
	if err != nil {
		log.Fatal("listen error:", err)
	}
	defer l.Close()

	for {
		// Accept incoming client request and parse data
		conn, err := l.Accept()
		if err != nil {
			log.Fatal("accept error:", err)
		}

		go server(conn)
	}
}
