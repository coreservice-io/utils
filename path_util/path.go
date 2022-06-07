package path_util

import (
	"os"
	"path/filepath"
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

// input a relative path or a absolute path
// return (abs_path,exist,err)
// if op-system file error then nothing meaningful
func SmartPathExist(abs_or_rel_path string) (string, bool, error) {

	//check abs path or relative path
	is_abs := filepath.IsAbs(abs_or_rel_path)
	if is_abs {
		//check if abs path exist
		abs_exist, err := AbsPathExist(abs_or_rel_path)
		if err != nil {
			return "", false, err
		}
		if abs_exist {
			return abs_or_rel_path, true, nil
		} else {
			return abs_or_rel_path, false, nil
		}
	} else {

		//rel to exe
		exist, err := AbsPathExist(ExE_Path(abs_or_rel_path))
		if err != nil {
			return "", false, err
		}
		if exist {
			return ExE_Path(abs_or_rel_path), true, nil
		}

		////////for debug mode direct run go run ./
		/////////if user run from root as working directory
		currDir, err := os.Getwd()
		if err != nil {
			return "", false, err
		}

		w_path := filepath.Join(currDir, abs_or_rel_path)
		w_p_exist, err := AbsPathExist(w_path)
		if err != nil {
			return "", false, err
		}
		if w_p_exist {
			return w_path, true, nil
		}
		//////////////////////////////////
		return ExE_Path(abs_or_rel_path), false, nil
	}
}
