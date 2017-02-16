package common

import (
	"github.com/wtsi-hgi/packer/helper/communicator"
	"github.com/wtsi-hgi/packer/template/interpolate"
)

type SSHConfig struct {
	Comm communicator.Config `mapstructure:",squash"`
}

func (c *SSHConfig) Prepare(ctx *interpolate.Context) []error {
	return c.Comm.Prepare(ctx)
}
