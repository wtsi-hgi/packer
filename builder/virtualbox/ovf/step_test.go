package ovf

import (
	"bytes"
	"github.com/mitchellh/multistep"
	vboxcommon "github.com/wtsi-hgi/packer/builder/virtualbox/common"
	"github.com/wtsi-hgi/packer/packer"
	"testing"
)

func testState(t *testing.T) multistep.StateBag {
	state := new(multistep.BasicStateBag)
	state.Put("driver", new(vboxcommon.DriverMock))
	state.Put("ui", &packer.BasicUi{
		Reader: new(bytes.Buffer),
		Writer: new(bytes.Buffer),
	})
	return state
}
