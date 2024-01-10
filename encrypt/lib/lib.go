package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"errors"
	"io"
	"sync"
)

var cache = struct {
	sync.RWMutex
	m map[string][]byte
}{m: make(map[string][]byte)}

func ConcurrentEncrypt(plaintexts [][]byte, key []byte) ([][]byte, error) {
	ciphertexts := make([][]byte, len(plaintexts))

	var wg sync.WaitGroup
	for i, plaintext := range plaintexts {
		wg.Add(1)
		go func(i int, plaintext []byte) {
			defer wg.Done()
			ciphertext, err := AdvancedEncrypt(plaintext, key)
			if err != nil {
				return
			}
			ciphertexts[i] = ciphertext
		}(i, plaintext)
	}
	wg.Wait()

	return ciphertexts, nil
}

func ConcurrentDecrypt(ciphertexts [][]byte, key []byte) ([][]byte, error) {
	plaintexts := make([][]byte, len(ciphertexts))

	var wg sync.WaitGroup
	for i, ciphertext := range ciphertexts {
		wg.Add(1)
		go func(i int, ciphertext []byte) {
			defer wg.Done()
			plaintext, err := AdvancedDecrypt(ciphertext, key)
			if err != nil {
				return
			}
			plaintexts[i] = plaintext
		}(i, ciphertext)
	}
	wg.Wait()

	return plaintexts, nil
}

func AdvancedEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	cache.RLock()
	ciphertext, ok := cache.m[string(plaintext)]
	cache.RUnlock()

	if ok {
		return ciphertext, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext = make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	cache.Lock()
	cache.m[string(plaintext)] = ciphertext
	cache.Unlock()

	return ciphertext, nil
}

func AdvancedDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	cache.RLock()
	plaintext, ok := cache.m[string(ciphertext)]
	cache.RUnlock()

	if ok {
		return plaintext, nil
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	cache.Lock()
	cache.m[string(ciphertext)] = ciphertext
	cache.Unlock()

	return ciphertext, nil
}

func Encrypt(plaintext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	ciphertext := make([]byte, aes.BlockSize+len(plaintext))
	iv := ciphertext[:aes.BlockSize]
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return nil, err
	}

	stream := cipher.NewCFBEncrypter(block, iv)
	stream.XORKeyStream(ciphertext[aes.BlockSize:], plaintext)

	return ciphertext, nil
}

func Decrypt(ciphertext []byte, key []byte) ([]byte, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	if len(ciphertext) < aes.BlockSize {
		return nil, errors.New("ciphertext too short")
	}

	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(ciphertext, ciphertext)

	return ciphertext, nil
}
