package common

import (
	"github.com/wtsi-hgi/packer/template/interpolate"
)

type VMXConfig struct {
	VMXData     map[string]string `mapstructure:"vmx_data"`
	VMXDataPost map[string]string `mapstructure:"vmx_data_post"`
}

func (c *VMXConfig) Prepare(ctx *interpolate.Context) []error {
	return nil
}
