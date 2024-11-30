package arg

import (
	"os"
)

var (
	InstanceDescription string
	InstanceName        string
	InstanceHost        string

	UserHandle  string
)

func init() {
	InstanceDescription = os.Getenv("MICRODON_INSTANCE_DESCRIPTION")
	InstanceName        = os.Getenv("MICRODON_INSTANCE_NAME")
	InstanceHost        = os.Getenv("MICRODON_INSTANCE_HOST")

	UserHandle  = os.Getenv("MICRODON_USER_HANDLE")
}
