package path_util

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"strings"
)

var exePath string

func ExEPathPrintln() {
	fmt.Println("EXE abs path:", exePath)
}

//return absolute path given relative path to the executable file folder
func GetAbsPath(relpath string) string {
	return filepath.Join(exePath, relpath)
}

// exists returns whether the given file or directory exists, when err exist not sure
func AbsPathExist(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
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
