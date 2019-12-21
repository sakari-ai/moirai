package namespace

import (
	"fmt"
	"os"

	"github.com/sakari-ai/moirai/config/env"
)

const (
	namespaceFmt = "github.com/sakari-ai/moirai/%s/services/%s"
)

type Namespace string

func (n Namespace) String() string {
	return string(n)
}

func FromString(value string) Namespace {
	return Namespace(value)
}

func FromMode(service string) Namespace {
	mode := os.Getenv(env.KeyConfigMode)
	if mode == "" {
		mode = "prod"
	}
	return FromString(fmt.Sprintf(namespaceFmt, mode, service))
}
