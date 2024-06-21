package main

import (
	"fmt"
	"os"
	"path/filepath"
	"sync"
)

func main() {
	folders := []string{"1", "2", "3"}

	currentPath, err := os.Getwd()
	if err != nil {
		fmt.Println(err)
		return
	}

	testPath := filepath.Join(currentPath, ".test")

	err = os.Mkdir(testPath, 0755)
	if err != nil {
		fmt.Println(err)
		return
	}

	wg := sync.WaitGroup{}

	err = setup(testPath, folders, &wg)
	if err != nil {
		fmt.Println(err)
		return
	}

	wg.Wait()

	fmt.Println("Folder structure created successfully.")
}

func setup(path string, folders []string, wg *sync.WaitGroup) error {
	defer wg.Done()

	for _, folder := range folders {
		newPath := filepath.Join(path, folder)

		err := os.Mkdir(newPath, 0755)
		if err != nil {
			return err
		}

		wg.Add(1)
		go setup(newPath, folders[:len(folders)-1], wg)
	}

	filePath := filepath.Join(path, "example.txt")

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString("this is an example file")
	if err != nil {
		return err
	}

	return nil
}

func cleanup(path string) error {
	err := os.RemoveAll(path)
	if err != nil {
		return fmt.Errorf("error deleting folder %s: %w", path, err)
	}
	return nil
}
