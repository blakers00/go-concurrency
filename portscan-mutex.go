package main

import (
	"fmt"
	"net"
	"strconv"
	"sync"
)

type Portscan struct {
	host    string
	port    []int
	kind    string
	timeout int
}

func (P *Portscan) solve(wg *sync.WaitGroup, mg *sync.Mutex, k int) {
	kind := P.kind
	host := P.host
	mg.Lock()
	data := k
	mg.Unlock()

	defer wg.Done()
	conn, _ := net.Dial(kind, host+":"+strconv.Itoa(data))

	if conn != nil {
		fmt.Printf("Port: %v open\n", data)
	} else {
		fmt.Printf("Port %v closed\n", data)
	}
}

func main() {
	i := Portscan{host: "172.24.144.1", port: []int{80, 443, 8080, 53, 21, 22, 3389}, kind: "tcp", timeout: 5}
	var wg sync.WaitGroup
	var mg sync.Mutex
	wg.Add(len(i.port))
	for _, k := range i.port {
		go i.solve(&wg, &mg, k)
	}
	wg.Wait()
}
