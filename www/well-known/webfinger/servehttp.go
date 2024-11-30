package verboten

import (
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

func serveWebFinger(resource string, rels ...string) ([]byte, error) {

	var response = webfinger.Response {
		Subject: opt.Something(resource),
	}

	return response.MarshalJSON()
}
