package verboten

import (
	"encoding/json"
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-webfinger"

	"github.com/reiver/microdon/srv/http"
)

const path string = webfinger.DefaultPath

func init() {
var webFingerHandler webfinger.Handler = webfinger.HandlerFunc(serveWebFinger)

	var handler http.Handler = webfinger.HTTPHandler(webFingerHandler)

	err := httpsrv.Mux.HandlePath(handler, path)
	if nil != err {
		err = erorr.Errorf("problem registered HTTP-handler for path %q: %w", path, err)
		panic(err)
	}
}

func serveWebFinger(responseWriter http.ResponseWriter, resource string, rels ...string) {

			
//@TODO
			

	var response = webfinger.Response {
		Subject: opt.Something(resource),
	}

	err := json.NewEncoder(responseWriter).Encode(response)
	_ = err
}
