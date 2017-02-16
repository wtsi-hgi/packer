package null

import (
	"fmt"

	"github.com/wtsi-hgi/packer/common"
	"github.com/wtsi-hgi/packer/helper/communicator"
	"github.com/wtsi-hgi/packer/helper/config"
	"github.com/wtsi-hgi/packer/packer"
	"github.com/wtsi-hgi/packer/template/interpolate"
)

type Config struct {
	common.PackerConfig `mapstructure:",squash"`

	CommConfig communicator.Config `mapstructure:",squash"`
}

func NewConfig(raws ...interface{}) (*Config, []string, error) {
	var c Config

	err := config.Decode(&c, &config.DecodeOpts{
		Interpolate:       true,
		InterpolateFilter: &interpolate.RenderFilter{},
	}, raws...)
	if err != nil {
		return nil, nil, err
	}

	var errs *packer.MultiError
	if es := c.CommConfig.Prepare(nil); len(es) > 0 {
		errs = packer.MultiErrorAppend(errs, es...)
	}
	if c.CommConfig.Host() == "" {
		errs = packer.MultiErrorAppend(errs,
			fmt.Errorf("a Host must be specified, please reference your communicator documentation"))
	}

	if c.CommConfig.User() == "" {
		errs = packer.MultiErrorAppend(errs,
			fmt.Errorf("a Username must be specified, please reference your communicator documentation"))
	}

	if c.CommConfig.Password() == "" && c.CommConfig.SSHPrivateKey == "" {
		errs = packer.MultiErrorAppend(errs,
			fmt.Errorf("one authentication method must be specified, please reference your communicator documentation"))
	}

	if c.CommConfig.SSHPassword != "" && c.CommConfig.SSHPrivateKey != "" {
		errs = packer.MultiErrorAppend(errs,
			fmt.Errorf("only one of ssh_password and ssh_private_key_file must be specified"))
	}

	if errs != nil && len(errs.Errors) > 0 {
		return nil, nil, errs
	}

	return &c, nil, nil
}
