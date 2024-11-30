package arg

import (
	"os"
)

var (
	InstanceName        string
	InstanceDescription string
)

func init() {
	InstanceName        = os.Getenv("MICRODON_INSTANCE_NAME")
	InstanceDescription = os.Getenv("MICRODON_INSTANCE_DESCRIPTION")
}
