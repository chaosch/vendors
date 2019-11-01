package common

import (
	"os"
)

func PathExists(path string) (bool, error) {
	_, err := os.Stat(path)
	if err == nil {
		return true, nil
	}
	if os.IsNotExist(err) {
		return false, nil
	}
	return false, err
}

func CreatePath(path string) error {
	_dir := path
	exist, err := PathExists(_dir)
	if err != nil {
		return err
	}

	if exist {
		return nil
	} else {
		err := os.Mkdir(_dir, os.ModePerm)
		return err
	}
}
