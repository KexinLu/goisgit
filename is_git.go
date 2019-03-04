package is_git

import (
	"errors"
	"github.com/spf13/afero"
	"fmt"
	"os/exec"
	"strings"
	"bytes"
)

const NOT_A_GIT = `not a git repository`
const IS_A_GIT = `true`

var ErrNotDir = errors.New(`target path is not a directory`)
var ErrFailToValidateDir = errors.New(`failed to validate is directory`)
var ErrFailToExecGitCmd = errors.New(`failed to execute git command`)

var GitCmdStr = `git`

var fs afero.Fs
func init() {
	fs = afero.NewOsFs()
}

func IsGitDir(path string) (bool, error) {
	return IsGitDirFs(&fs, path)
}

func IsGitDirFs(fs *afero.Fs, path string) (bool, error) {
	if is, err := afero.IsDir(*fs, path); err != nil {
		return false, ErrFailToValidateDir
	} else if !is {
		return false, ErrNotDir
	}

	cmd := exec.Command(GitCmdStr, "rev-parse", "--is-inside-work-tree")

	errBuf := bytes.NewBuffer(make([]byte, 100))
	cmd.Dir = path
	cmd.Stderr = errBuf

	if out, err := cmd.Output(); err != nil {
		errStr := errBuf.String()
		if strings.Contains(errStr, NOT_A_GIT) {
			return false, nil
		}

		return false, fmt.Errorf("%s: %s\n", ErrFailToExecGitCmd.Error(), err.Error())
	} else if strings.Contains(string(out), IS_A_GIT) {
		return true, nil
	}

	return false, nil
}
