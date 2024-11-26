package cfg

import (
	"github.com/reiver/go-opt"
)

const (
	SoftwareName    string = "microdon"
	SoftwareVersion string = "0.0.0"
)

var (
	SoftwareRepository opt.Optional[string] = opt.Something("https://github.com/reiver/microdon")
	SoftwareHomePage   opt.Optional[string] = opt.Nothing[string]()
)
