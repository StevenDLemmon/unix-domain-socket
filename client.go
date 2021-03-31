package main

import (
	"log"
	"net"
	"os/exec"
)

func main() {
	c, err := net.Dial("unix", "/tmp/memory.sock")
	if err != nil {
		log.Fatal(err)
	}
	defer c.Close()

	// This linux command prints out all PIDs with the memory they are currently using (ps -No pid:1,size:1),
	// removes the first line of the output, which are the headers (sed '1d'), and then sorts the second 
	// column in descending order (sort -r -nk +2)
	out, err := exec.Command("bash", "-c", "ps -No pid:1,size:1 | sed '1d' | sort -r -nk +2").Output()

	if err != nil {
		log.Fatal(err)
	}

	// Send the output to the server to be parse
	_, err = c.Write(out)
	if err != nil {
		log.Fatal("write error:", err)
	}
}
