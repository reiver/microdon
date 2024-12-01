package verboten

import (
	"fmt"
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-errhttp"
	"github.com/reiver/go-opt"
	"github.com/reiver/go-webfinger"

	"github.com/reiver/microdon/arg"
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

	var userHandle   string = arg.UserHandle
	var instanceHost string = arg.InstanceHost

	if "" == userHandle {
		return nil, errhttp.Return(http.StatusNotFound)
	}

	if "" == instanceHost {
		return nil, errhttp.Return(http.StatusNotFound)
	}

	var acctURI string = fmt.Sprintf("acct:%s@%s", userHandle, instanceHost)

	if acctURI != resource {
		return nil, errhttp.Return(http.StatusNotFound)
	}

	var profileLink string = fmt.Sprintf("https://%s/", instanceHost)

	var activityLink string = fmt.Sprintf("%s-/rel/self.activity", profileLink)

	{
		var response = webfinger.Response {
			Subject: opt.Something(resource),
			Links: []webfinger.Link{
				webfinger.Link{
					Rel: opt.Something("http://webfinger.net/rel/profile-page"),
					Type: opt.Something("text/html"),
					HRef: opt.Something(profileLink),
				},
				webfinger.Link{
					Rel: opt.Something("self"),
					Type: opt.Something("application/activity+json"),
					HRef: opt.Something(activityLink),
				},
			},
		}

		return response.MarshalJSON()
	}
}
