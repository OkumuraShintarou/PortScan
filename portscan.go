package main

import (
	"fmt"
	"net"
	"strconv"
	"time"
)

func main() {

	ipaddr := "95.211.187.146"
	ports := []int{21, 22, 23, 25, 80, 110, 111, 443, 445, 3306, 3389, 5900, 5901, 8080, 9050, 9051}

	for _, port := range ports {
		resp := check(ipaddr, port)

		if resp == 0 {
			fmt.Printf("open: %d\n", port)
		}
	}

}

func check(host string, port int) int {

	target := host + ":" + strconv.Itoa(port)

	server, err := net.DialTimeout("tcp", target, 3*time.Second)

	if err != nil {
		return 1
	}

	server.Close()

	return 0

}
