package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
)

var (
	block cipher.Block
	nonce []byte
)

// KeyGenAndInit generates a 32-bytes secret key and initializes the AES cipher,
// and returns the secret key and nonce or an error if any operation fails.
func KeyGenAndInit() ([]byte, []byte, error) {
	secretKey, err := genSecretKey()
	if err != nil {
		return nil, nil, err
	}

	err = Init(secretKey)
	if err != nil {
		return nil, nil, err
	}

	return secretKey, nonce, nil
}

// Init generates a 16-byte nonce and create the AES cipher using the provided secret key.
func Init(secretKey []byte) (err error) {
	nonce, err = genNonce()
	if err != nil {
		return err
	}

	block, err = aes.NewCipher(secretKey)
	return err
}

// GenSecretKey generates a random 32-byte (128 bits) key for AES-256 encryption.
// It returns the key or an error if the key generation fails.
func genSecretKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// GenNonce generates a random 16-byte nonce for cryptographic operations using AES block size.
// It returns the nonce and any error encountered during the generation.
func genNonce() ([]byte, error) {
	nc := make([]byte, aes.BlockSize)

	_, err := rand.Read(nc)
	if err != nil {
		return nil, err
	}

	return nc, nil
}

// Encrypt encrypts the given plaintext using AES-256 CTR mode.
func Encrypt(plaintext []byte) []byte {
	plaintext = pkcs7pad(plaintext)
	ciphertext := make([]byte, len(plaintext))

	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext
}

// Decrypt decrypts the given ciphertext using AES-256 CTR mode.
func Decrypt(ciphertext []byte) []byte {

	plaintext := make([]byte, len(ciphertext))

	stream := cipher.NewCTR(block, nonce)
	stream.XORKeyStream(plaintext, ciphertext)

	plaintext = pkcs7strip(plaintext)
	return plaintext
}

// pkcs7pad add pkcs7 padding
func pkcs7pad(data []byte) []byte {
	padLen := aes.BlockSize - len(data)%aes.BlockSize
	padding := bytes.Repeat([]byte{byte(padLen)}, padLen)

	return append(data, padding...)
}

// pkcs7strip remove pkcs7 padding
func pkcs7strip(data []byte) []byte {
	length := len(data)
	padLen := int(data[length-1])

	if padLen == 0 {
		return data
	}
	return data[:length-padLen]
}
