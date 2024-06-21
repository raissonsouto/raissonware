package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

// GenSecretKey generates a random 32-byte (128 bits) key for AES-256 encryption.
// It returns the key or an error if the key generation fails.
func GenSecretKey() ([]byte, error) {
	key := make([]byte, 32)

	_, err := rand.Read(key)
	if err != nil {
		return nil, err
	}

	return key, nil
}

// GenNonce generates a random 16-byte nonce for cryptographic operations using AES block size.
// It returns the nonce and any error encountered during the generation.
func GenNonce() ([]byte, error) {
	nonce := make([]byte, aes.BlockSize)

	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	return nonce, nil
}

// Encrypt encrypts the given plaintext using AES-256 CTR mode.
// It returns the ciphertext or an error if the encryption fails.
func Encrypt(plaintext []byte, key []byte, nonce []byte) ([]byte, error) {
	aesMode, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	paddingSize := aes.BlockSize - (len(plaintext) % aes.BlockSize)

	if paddingSize != 0 {
		pad := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
		plaintext = append(plaintext, pad...)
	}

	stream := cipher.NewCTR(aesMode, nonce)

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

// Decrypt decrypts the given ciphertext using AES-256 CTR mode.
// It returns the plaintext or an error if the decryption fails.
func Decrypt(ciphertext []byte, key []byte, nonce []byte) ([]byte, error) {
	aesMode, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext shorter than block size")
	}

	stream := cipher.NewCTR(aesMode, nonce)

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
