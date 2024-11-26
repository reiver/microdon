package main

import (
	"fmt"
	"net/http"

	_ "github.com/reiver/microdon/www"
)

func webserver(tcpPort uint16, handler http.Handler) {

	var tcpAddress string = fmt.Sprintf(":%d", tcpPort)

	http.ListenAndServe(tcpAddress, handler)
}
