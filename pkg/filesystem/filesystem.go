package filesystem

import (
	"os"
	"path/filepath"
	"sync"
)

// FindFiles recursively search for files and executes the given function
func FindFiles(path string, exec func([]byte) ([]byte, error), wg *sync.WaitGroup, chanErr chan error) {
	defer wg.Done()

	files, err := os.ReadDir(path)
	if err != nil {
		chanErr <- err
		return
	}

	for _, file := range files {
		subPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			wg.Add(1)
			go FindFiles(subPath, exec, wg, chanErr)
			continue
		}

		alterFile(subPath, exec, chanErr)
	}
}

func alterFile(path string, exec func([]byte) ([]byte, error), chanErr chan error) {

	fileInfo, err := os.Stat(path)
	if err != nil {
		chanErr <- err
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		chanErr <- err
		return
	}

	newData, err := exec(data)
	if err != nil {
		chanErr <- err
		return
	}

	err = os.WriteFile(path, newData, fileInfo.Mode())
	if err != nil {
		chanErr <- err
		return
	}
}
