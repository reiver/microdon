package nodeinfo1dot1

import (
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-nodeinfo/1.1"

	"github.com/reiver/microdon/cfg"
	microdonNodeInfo "github.com/reiver/microdon/lib/nodeinfo"
	"github.com/reiver/microdon/srv/http"
)

const Path string = "/-/nodeinfo/1.1"

var paths = []string{
	Path,                     // microdon style
	"/nodeinfo/1.1",          // mastodon & firefish style
	"/nodeinfo/1.1.json",     // akkoma style
	"/api/nodeinfo/1.1.json", // pixelfed style
}

var handler http.Handler = nodeinfo.NodeInfo{
	Software:  microdonNodeInfo.Software1,
	Protocols: nodeinfo.Protocols{
		Inbound:  microdonNodeInfo.Protocols,
		Outbound: microdonNodeInfo.Protocols,
	},
	Services: nodeinfo.Services{
		Inbound:  []string{},
		Outbound: []string{},
	},
	OpenRegistrations: false,
	Usage: nodeinfo.Usage{
		Users: nodeinfo.Users{
			Total:          cfg.UsersTotal,
			ActiveHalfYear: cfg.UsersActiveHalfYear,
			ActiveMonth:    cfg.UsersActiveMonth,
		},
	},
	MetaData: nil,
}

func init() {
	for _, p := range paths {
		err := httpsrv.Mux.HandlePath(handler, p)
		if nil != err {
			err = erorr.Errorf("problem registered HTTP-handler for path %q: %w", p, err)
			panic(err)
		}
	}
}
