package util

import (
	"os"
	"path/filepath"
	"strings"
)

type Path struct{}

func NewPath() *Path {
	return &Path{}
}

func (p *Path) Code() (string, error) {
	path, err := os.Executable()
	if err != nil {
		return "", err
	}
	return filepath.Abs(path)
}

func (p *Path) Work() (string, error) {
	path, err := os.Getwd()
	if err != nil {
		return "", err
	}
	return filepath.Abs(path)
}

func (p *Path) Name(path string) string {
	file := filepath.Base(path)
	suffix := filepath.Ext(file)
	return strings.TrimSuffix(file, suffix)
}

func (p *Path) Exist(path string) bool {
	_, err := os.Stat(path)
	if err != nil {
		return os.IsExist(err)
	}
	return true
}
