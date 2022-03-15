package path_util

import (
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var exePath string

func ExE_PathStr() string {
	return exePath
}

//return absolute path given relative path to the executable file folder
func ExE_Path(relpath string) string {
	return filepath.Join(exePath, relpath)
}

func init() {
	file, err := exec.LookPath(os.Args[0])
	if err != nil {
		panic(err.Error())
	}
	runPath, err := filepath.Abs(file)
	if err != nil {
		panic(err.Error())
	}
	index := strings.LastIndex(runPath, string(os.PathSeparator))
	exePath = runPath[:index]
}
