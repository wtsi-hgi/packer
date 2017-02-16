package triton

import (
	"github.com/wtsi-hgi/packer/common"
	"github.com/wtsi-hgi/packer/helper/communicator"
	"github.com/wtsi-hgi/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`
	AccessConfig        `mapstructure:",squash"`
	SourceMachineConfig `mapstructure:",squash"`
	TargetImageConfig   `mapstructure:",squash"`

	Comm communicator.Config `mapstructure:",squash"`

	ctx interpolate.Context
}
