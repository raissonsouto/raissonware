package cryptography

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
)

var stream cipher.Stream

func KeyGenAndInit() ([]byte, []byte, error) {
	secretKey, err := genSecretKey()
	if err != nil {
		return nil, nil, err
	}

	nonce, err := genNonce()
	if err != nil {
		return nil, nil, err
	}

	err = Init(secretKey, nonce)
	if err != nil {
		return nil, nil, err
	}

	return secretKey, nonce, nil
}

func Init(secretKey []byte, nonce []byte) error {
	aesMode, err := aes.NewCipher(secretKey)
	if err != nil {
		return err
	}

	stream = cipher.NewCTR(aesMode, nonce)
	return nil
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
	nonce := make([]byte, aes.BlockSize)

	_, err := rand.Read(nonce)
	if err != nil {
		return nil, err
	}

	return nonce, nil
}

// Encrypt encrypts the given plaintext using AES-256 CTR mode.
// It returns the ciphertext or an error if the encryption fails.
func Encrypt(plaintext []byte) ([]byte, error) {
	paddingSize := aes.BlockSize - (len(plaintext) % aes.BlockSize)

	if paddingSize != 0 {
		pad := bytes.Repeat([]byte{byte(paddingSize)}, paddingSize)
		plaintext = append(plaintext, pad...)
	}

	ciphertext := make([]byte, len(plaintext))
	stream.XORKeyStream(ciphertext, plaintext)

	return ciphertext, nil
}

// Decrypt decrypts the given ciphertext using AES-256 CTR mode.
// It returns the plaintext or an error if the decryption fails.
func Decrypt(ciphertext []byte) ([]byte, error) {
	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext shorter than block size")
	}

	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}
