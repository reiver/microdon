package cfg

import (
	"github.com/reiver/go-opt"
)

var (
	UsersTotal          opt.Optional[string] = opt.Something("1")
	UsersActiveHalfYear opt.Optional[string] = opt.Something("1")
	UsersActiveMonth    opt.Optional[string] = opt.Something("1")
	UsersActiveWeek     opt.Optional[string] = opt.Something("1")
)
