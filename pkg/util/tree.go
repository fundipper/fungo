package util

import (
	"bufio"
	"io"
	"os"
	"path/filepath"
	"strings"
)

type Tree struct{}

func NewTree() *Tree {
	return &Tree{}
}

func (t *Tree) ReadDir(name string) ([]string, error) {
	result := []string{}
	err := filepath.Walk(name, func(path string, info os.FileInfo, err error) error {
		if info != nil {
			if info.IsDir() {
				return nil
			}

			result = append(result, path)
		}
		return nil
	})
	return result, err
}

func (t *Tree) MkDir(path string) error {
	_, err := os.Stat(path)
	if err != nil && os.IsExist(err) {
		return err
	}
	return os.MkdirAll(path, os.ModePerm)
}

func (t *Tree) ReadFile(name string) ([]byte, error) {
	f, err := os.Open(name)
	if err != nil {
		return nil, err
	}
	defer f.Close()
	reader := bufio.NewReader(f)

	var chunk []byte
	block := make([]byte, 1024)
	for {
		n, err := reader.Read(block)
		if err != nil && err != io.EOF {
			return nil, err
		}
		if n == 0 {
			break
		}
		chunk = append(chunk, block[:n]...)
	}
	return chunk, nil
}

func (t *Tree) WriteFile(name string, data []byte) (int, error) {
	_ = NewTree().MkDir(filepath.Dir(name))
	f, err := os.OpenFile(strings.ToLower(name), os.O_CREATE|os.O_RDWR|os.O_TRUNC, os.ModePerm)
	if err != nil {
		return 0, err
	}
	defer f.Close()

	writer := bufio.NewWriter(f)
	nn, err := writer.Write(data)
	if err != nil {
		return 0, err
	}
	err = writer.Flush()
	if err != nil {
		return 0, err
	}
	return nn, nil
}
