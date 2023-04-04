package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"net"
	"strconv"
	"strings"
	"time"
)

const MIN = 1
const MAX = 100

var	serverAddr string
var serverPort string

func init() {
	serverAddr = "0.0.0.0"
	serverPort = ":2244"
}

func random() int {
	return rand.Intn(MAX-MIN) + MIN
}

func handleConnection(c net.Conn) {
	fmt.Printf("Serving %s\n", c.RemoteAddr().String())
	for {
		netData, err := bufio.NewReader(c).ReadString('\n')
		fmt.Printf("Data: %s\n", netData)
		if err != nil {
			fmt.Println(err)
			return
		}

		temp := strings.TrimSpace(string(netData))
		if temp == "STOP" {
			fmt.Fprint(c, "you said %s so, bye\n", temp)
			break
		}

		result := strconv.Itoa(random()) + "\n"
		c.Write([]byte(string(result)))
	}
	c.Close()
}


// NOTE: Read this later https://gist.github.com/spikebike/2232102 (For TLS Conn)
func main() {
	l, err := net.Listen("tcp4", serverPort)
	fmt.Printf("Secretary listening to %s:%s\n", serverAddr, serverPort)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		go handleConnection(c)
	}
}
