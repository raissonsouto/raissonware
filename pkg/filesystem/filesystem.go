package filesystem

import (
	"os"
	"path/filepath"
	"sync"
)

// FindFiles recursively search for files and executes the given function
func FindFiles(path string, exec func([]byte) []byte, wg *sync.WaitGroup, chanErr chan error) {
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

// alterFile reads the contents of a file at the specified path,
// executes the provided function on the data, writes back the modified data,
// and handles any errors via the provided error channel.
func alterFile(path string, exec func([]byte) []byte, chanErr chan error) {

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

	newData := exec(data)

	err = os.WriteFile(path, newData, fileInfo.Mode())
	if err != nil {
		chanErr <- err
		return
	}
}
