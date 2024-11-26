package main

import (
	"fmt"
	"net/http"
	"os"

	"github.com/reiver/microdon/srv/http"
)

func main() {

	fmt.Println("microdon 🐟")

	var tcpPort uint16
	{
		var err error

		tcpPort, err = tcpport()
		if nil != err {
			fmt.Fprintf(os.Stderr, "ERROR: problem getting TCP-port (used by the web-server) from environent-variable 'PORT' (which should be a 16-bit unsigned number): %s\n", err)
			os.Exit(1)
			return
		}
	}

	{
		var handler http.Handler = &httpsrv.Mux

		webserver(tcpPort, handler)
	}
}
