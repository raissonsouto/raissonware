package main

import (
	"os"
	"path/filepath"
)

// createDotTest creates the root testing environment folder and return its path
func createDotTest(folders []string) (string, error) {
	currentPath, err := os.Getwd()
	if err != nil {
		return "", err
	}

	testingPath := filepath.Join(currentPath, testingFolder)
	if err = os.Mkdir(testingPath, 0755); err != nil {
		return "", err
	}

	chanErr := make(chan error)

	wg.Add(1)
	go createSubFolders(testingPath, folders, chanErr)

	wg.Wait()

	if len(chanErr) != 0 {
		for err = range chanErr {
			return "", err
		}
	}

	err = createExampleFile(testingPath, fileText)
	if err != nil {
		return "", err
	}

	return testingPath, nil
}

// createSubFolders creates subfolders and example files recursively
func createSubFolders(path string, folders []string, chanErr chan error) {
	defer wg.Done()

	for _, folder := range folders {
		subPath := filepath.Join(path, folder)

		err := os.Mkdir(subPath, 0755)
		if err != nil {
			chanErr <- err
			return
		}

		wg.Add(1)
		go createSubFolders(subPath, folders[:len(folders)-1], chanErr)

		err = createExampleFile(subPath, fileText)
		if err != nil {
			chanErr <- err
			return
		}
	}
}

// createExampleFile creates an example.txt file in the given path and writes the given text
func createExampleFile(path string, fileText string) (err error) {
	filePath := filepath.Join(path, exampleFileName)

	file, err := os.Create(filePath)
	if err != nil {
		return err
	}

	defer file.Close()

	_, err = file.WriteString(fileText)
	if err != nil {
		return err
	}

	return err
}
