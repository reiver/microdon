package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/shared"
	"github.com/reiver/go-opt"

	"github.com/reiver/microdon/cfg"
)

const softwareName       string = cfg.SoftwareName
const softwareVersion    string = cfg.SoftwareVersion
var   softwareRepository opt.Optional[string] = cfg.SoftwareRepository
var   softwareHomePage   opt.Optional[string] = cfg.SoftwareHomePage

var Software1 = shared.Software1{
	Name:    softwareName,
	Version: softwareVersion,
}

var Software2Dot1 = shared.Software2Dot1{
	Name:       softwareName,
	Version:    softwareVersion,
	Repository: softwareRepository,
	HomePage:   softwareHomePage,
}
