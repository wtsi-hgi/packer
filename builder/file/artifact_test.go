package file

import (
	"testing"

	"github.com/wtsi-hgi/packer/packer"
)

func TestNullArtifact(t *testing.T) {
	var _ packer.Artifact = new(FileArtifact)
}
