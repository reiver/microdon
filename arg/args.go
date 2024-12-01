package arg

import (
	"os"
)

var (
	InstanceDescription string
	InstanceHost        string
	InstanceName        string

	UserHandle  string
)

func init() {
	InstanceDescription = os.Getenv("MICRODON_INSTANCE_DESCRIPTION")
	InstanceHost        = os.Getenv("MICRODON_INSTANCE_HOST")
	InstanceName        = os.Getenv("MICRODON_INSTANCE_NAME")

	UserHandle  = os.Getenv("MICRODON_USER_HANDLE")
}
