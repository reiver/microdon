package nodeinfo

import (
	"github.com/reiver/go-nodeinfo/shared"
	"github.com/reiver/go-opt"

	"github.com/reiver/microdon/arg"
)

var instanceName        string = arg.InstanceName
var instanceDescription string = arg.InstanceDescription

var Instance = shared.Instance2Dot2{}

func init() {
	if "" != instanceName {
		Instance.Name = opt.Something(instanceName)
	}
	if "" != instanceDescription {
		Instance.Description = opt.Something(instanceDescription)
	}
}
