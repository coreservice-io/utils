package path_util

import (
	"errors"
	"os"
	"path/filepath"
	//"github.com/labstack/echo/v4/middleware"
)

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

//return a abs folder/file path ,
// if no such file or folder then err is not nil
func SmartExistPath(relpath string) (string, error) {

	exist, _ := AbsPathExist(ExE_Path(relpath))
	if exist {
		return ExE_Path(relpath), nil
	}

	//if user run from root as working directory
	currDir, err := os.Getwd()
	if err != nil {
		return "", err
	}
	w_path := filepath.Join(currDir, relpath)
	w_p_exist, _ := AbsPathExist(w_path)
	if !w_p_exist {
		return "", errors.New("path not exist")
	}
	return w_path, nil
}
