package session

import (
	"os"
	"net"
	"time"
	"math/rand"

	"secretary/alpha/utils"
	"secretary/alpha/internal"
)

func ListenAndServeTCP() {
    sessionTCPPort := os.Getenv("SESSION_TCP_PORT")
    if sessionTCPPort == "" {
        sessionTCPPort = "4848"
    }
    utils.Logger("info", "TCP server listening on " + sessionTCPPort)
    l, err := net.Listen("tcp4", ":" + sessionTCPPort)
	if err != nil {
		utils.Logger("fatal", err.Error())
		return
	}
	defer l.Close()
	rand.Seed(time.Now().Unix())

	for {
		c, err := l.Accept()
		if err != nil {
		    utils.Logger("error", err.Error())
			return
		}
		go internal.HandleTCPConnection(c)
	}
}
