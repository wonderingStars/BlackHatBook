package main

import (
	"fmt"
	"net"
	"sort"
)

func worker(ports, results chan int, address string) {
	for p := range ports {
		address := fmt.Sprintf(address, p)
		conn, err := net.Dial("tcp", address)
		if err != nil {
			results <- 0
			continue
		}
		conn.Close()
		results <- p
	}
}

func ScanPortUseingWorkers(portMax int, address string) {
	ports := make(chan int, portMax)
	results := make(chan int)
	var openports []int
	for i := 0; i < cap(ports); i++ {
		go worker(ports, results, address)
	}
	go func() {
		for i := 1; i <= portMax; i++ {
			ports <- i
		}
	}()
	for i := 0; i < portMax; i++ {
		port := <-results
		if port != 0 {
			openports = append(openports, port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports)
	for _, port := range openports {
		fmt.Printf("%d open\n", port)
	}
}
