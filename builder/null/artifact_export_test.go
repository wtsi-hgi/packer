package null

import (
	"github.com/wtsi-hgi/packer/packer"
	"testing"
)

func TestNullArtifact(t *testing.T) {
	var _ packer.Artifact = new(NullArtifact)
}
