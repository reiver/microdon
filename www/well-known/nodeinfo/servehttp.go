package verboten

import (
	"fmt"
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-nodeinfo"
	"github.com/reiver/go-opt"

	"github.com/reiver/microdon/srv/http"
	"github.com/reiver/microdon/www/dash/nodeinfo/1.0"
	"github.com/reiver/microdon/www/dash/nodeinfo/1.1"
	"github.com/reiver/microdon/www/dash/nodeinfo/2.0"
	"github.com/reiver/microdon/www/dash/nodeinfo/2.1"
	"github.com/reiver/microdon/www/dash/nodeinfo/2.2"
)

const path string = "/.well-known/nodeinfo"

func init() {
	var handler http.Handler = http.HandlerFunc(serveHTTP)

	err := httpsrv.Mux.HandlePath(handler, path)
	if nil != err {
		err = erorr.Errorf("problem registered HTTP-handler for path %q: %w", path, err)
		panic(err)
	}
}

func serveHTTP(responseWriter http.ResponseWriter, request *http.Request) {
	if nil == responseWriter {
		return
	}
	if nil == request {
		var code int = http.StatusInternalServerError
		var text string = http.StatusText(code)

		http.Error(responseWriter, text, code)
		return
	}

	var host string = request.Host

	var nodeInfo2Dot2 string
	var nodeInfo2Dot1 string
	var nodeInfo2     string
	var nodeInfo1Dot1 string
	var nodeInfo1     string
	{
		var prefix string = fmt.Sprintf("http://%s", host)

		nodeInfo2Dot2 = prefix + nodeinfo2dot2.Path
		nodeInfo2Dot1 = prefix + nodeinfo2dot1.Path
		nodeInfo2     = prefix + nodeinfo2.Path
		nodeInfo1Dot1 = prefix + nodeinfo1dot1.Path
		nodeInfo1     = prefix + nodeinfo1.Path
	}


	var handler http.Handler = nodeinfo.WellKnown{
		NodeInfo2Dot2: opt.Something(nodeInfo2Dot2),
		NodeInfo2Dot1: opt.Something(nodeInfo2Dot1),
		NodeInfo2:     opt.Something(nodeInfo2),
		NodeInfo1Dot1: opt.Something(nodeInfo1Dot1),
		NodeInfo1:     opt.Something(nodeInfo1),
	}

	handler.ServeHTTP(responseWriter, request)
}
