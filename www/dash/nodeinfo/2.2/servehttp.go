package nodeinfo2dot2

import (
	"net/http"

	"github.com/reiver/go-erorr"
	"github.com/reiver/go-nodeinfo/2.2"

	"github.com/reiver/microdon/cfg"
	microdonNodeInfo "github.com/reiver/microdon/lib/nodeinfo"
	"github.com/reiver/microdon/srv/http"
)

const Path string = "/-/nodeinfo/2.2"

var paths = []string{
	Path,                     // microdon style
	"/nodeinfo/2.2",          // mastodon & firefish style
	"/nodeinfo/2.2.json",     // akkoma style
	"/api/nodeinfo/2.2.json", // pixelfed style
}

var handler http.Handler = nodeinfo.NodeInfo{
	Instance:  microdonNodeInfo.Instance,
	Software:  microdonNodeInfo.Software2Dot1,
	Protocols: microdonNodeInfo.Protocols,
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
			ActiveWeek:     cfg.UsersActiveWeek,
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
