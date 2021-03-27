package fs

import (
	"io/ioutil"
	"os"
	p "path"
)

// ReadText reads text file from the local file system.
func ReadText(path string) (string, error) {
	content, err := ioutil.ReadFile(path)
	if err != nil {
		return "", err
	}
	return string(content), nil
}

// WriteText write text to the local file system.
func WriteText(content string, path string) error {
	return WriteFile([]byte(content), path)
}

// WriteFile writes object to the local file system.
func WriteFile(content []byte, path string) error {
	if err := mkdirs(p.Dir(path)); err != nil {
		return err
	}
	return ioutil.WriteFile(path, content, 0644)
}

// mkdirs defines "mkdir -p ...".
func mkdirs(path string) error {
	return os.MkdirAll(path, 0755)
}
