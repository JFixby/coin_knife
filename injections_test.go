package coinknife

import (
	"testing"
	"path/filepath"
	"github.com/jfixby/pin"
	"github.com/jfixby/pin/fileops"
)

func TestListFiles(t *testing.T) {
	list := fileops.ListFiles(filepath.Join("", "code_injections"), FoldersOnlyAndIgnoreUnderscore, fileops.ALL_CHILDREN)

	pin.D("list", list)

}
