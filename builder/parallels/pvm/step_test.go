package pvm

import (
	"bytes"
	"testing"

	"github.com/mitchellh/multistep"
	parallelscommon "github.com/wtsi-hgi/packer/builder/parallels/common"
	"github.com/wtsi-hgi/packer/packer"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(parallelscommon.DriverMock))
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
