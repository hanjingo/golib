package file

import (
	"crypto/md5"
	"errors"
	"io"
	"os"
	"path/filepath"
)

func MD5(filePath string) ([]byte, error) {
	var result []byte
	file, err := os.Open(filePath)
	if err != nil {
		return result, err
	}
	defer file.Close()
	hash := md5.New()
	if _, err := io.Copy(hash, file); err != nil {
		return result, err
	}
	return hash.Sum(result), nil
}

func Exist(arg string) bool {
	_, err := os.Stat(arg) // get file info
	if err != nil {
		if os.IsExist(err) {
			return true
		}
		return false
	}
	return true
}

func Size(filePathName string) (SIZE, error) {
	if !Exist(filePathName) {
		return SIZE(0), errors.New("file not exist")
	}
	info, err := os.Stat(filePathName)
	if err != nil {
		return SIZE(0), errors.New("get file size fail")
	}
	return SIZE(info.Size()), nil
}

func FullName(filePathName string) string {
	_, name := filepath.Split(filePathName)
	return name
}

func NameAndExt(filePathName string) (string, string) {
	_, name := filepath.Split(filePathName)
	file_name := filepath.Base(name)
	file_type := filepath.Ext(filePathName)
	return file_name, file_type
}

func Create(filePathName string) (*os.File, error) {
	if !Exist(filePathName) {
		return os.Create(filePathName)
	}
	return nil, errors.New("file already exist")
}

func IsFile(file_path_name string) bool {
	s, err := os.Stat(file_path_name)
	return err == nil && !s.IsDir()
}

func IsDir(path string) bool {
	s, err := os.Stat(path)
	return err == nil && s.IsDir()
}
