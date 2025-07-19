package main

import (
	"net/http"
	"os"

	"github.com/reiver/microdon/srv/http"
	"github.com/reiver/microdon/srv/log"
)

func main() {

	log := logsrv.Prefix("main").Begin()
	defer log.End()

	log.Inform("microdon 🐟")

	var tcpPort uint16
	{
		var err error

		tcpPort, err = tcpport()
		if nil != err {
			log.Errorf("ERROR: problem getting TCP-port (used by the web-server) from environent-variable 'PORT' (which should be a 16-bit unsigned number): %s\n", err)
			os.Exit(1)
			return
		}
	}

	{
		var handler http.Handler = &httpsrv.Mux

		webserver(tcpPort, handler)
	}
}
