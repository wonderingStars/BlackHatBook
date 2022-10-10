package main

import (
	"fmt"
	"net"
)

func SimpleScan() {
	_, err := net.Dial("tcp", "scanme.nmap.org:80")
	if err == nil {
		fmt.Println("Connection successful")
	}
}
