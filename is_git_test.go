package is_git

import (
	"testing"
	"github.com/spf13/afero"
	"fmt"
)

func TestIsGitDirFsFails(t *testing.T) {
	fs := afero.NewMemMapFs()
	fs.Create("/file")
	table := []struct{
		path string
		err error
	}{
		{"/fi", ErrFailToValidateDir},
		{"/file", ErrNotDir},
	}

	for _, v := range table {
		_, err := IsGitDirFs(&fs, v.path)
		if err.Error() != v.err.Error() {
			fmt.Printf(`Expecting "%s" but got "%s"`, v.err.Error(), err.Error())
			t.Fail()
		}
	}
}

