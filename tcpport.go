package main

import (
	"os"
	"strconv"

	"github.com/reiver/go-erorr"
)

// tcpport returns the TCP-port that the HTTP server should use.
//
// It defaults to TCP-port 8080.
//
// But that can be overridden by a value set in the "PORT" environment variable.
func tcpport() (uint16, error) {
	var tcpPort uint16 = 8080

	{
		value := os.Getenv("PORT")
		if "" != value {
			u64, err := strconv.ParseUint(value, 10, 16)
			switch casted :=err.(type) {
			case *strconv.NumError:
				err = erorr.Errorf("problem converting the string %q into a number: %w", value, casted.Unwrap())
			}
			if nil != err {
				var nada uint16
				return nada, err
			}

			tcpPort = uint16(u64)
		}
	}

	return tcpPort, nil
}
