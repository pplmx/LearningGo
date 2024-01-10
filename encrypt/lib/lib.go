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

// ConcurrentEncrypt encrypts multiple plaintexts concurrently.
func ConcurrentEncrypt(plaintexts [][]byte, key []byte) ([][]byte, error) {
	// Create a slice to hold the ciphertexts.
	ciphertexts := make([][]byte, len(plaintexts))

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish.
	for i, plaintext := range plaintexts {
		wg.Add(1) // Add a count to the WaitGroup.
		go func(i int, plaintext []byte) {
			defer wg.Done()
			ciphertext, err := AdvancedEncrypt(plaintext, key)
			if err != nil {
				return
			}
			ciphertexts[i] = ciphertext // Store the ciphertext in the correct index.
		}(i, plaintext)
	}
	wg.Wait() // Wait for all goroutines to finish.

	return ciphertexts, nil
}

// ConcurrentDecrypt decrypts multiple ciphertexts concurrently.
func ConcurrentDecrypt(ciphertexts [][]byte, key []byte) ([][]byte, error) {
	// Create a slice to hold the plaintexts.
	plaintexts := make([][]byte, len(ciphertexts))

	var wg sync.WaitGroup // WaitGroup to wait for all goroutines to finish.
	for i, ciphertext := range ciphertexts {
		wg.Add(1) // Add a count to the WaitGroup.
		go func(i int, ciphertext []byte) {
			defer wg.Done()
			plaintext, err := AdvancedDecrypt(ciphertext, key)
			if err != nil {
				return
			}
			plaintexts[i] = plaintext // Store the plaintext in the correct index.
		}(i, ciphertext)
	}
	wg.Wait() // Wait for all goroutines to finish.

	return plaintexts, nil
}

// AdvancedEncrypt encrypts a plaintext and caches the result.
func AdvancedEncrypt(plaintext []byte, key []byte) ([]byte, error) {
	cache.RLock()                                // Acquire a read lock.
	ciphertext, ok := cache.m[string(plaintext)] // Check if the ciphertext is in the cache.
	cache.RUnlock()                              // Release the read lock.

	if ok {
		return ciphertext, nil // If the ciphertext is in the cache, return it.
	}

	// If the ciphertext is not in the cache, encrypt the plaintext.
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

	cache.Lock()                            // Acquire a write lock.
	cache.m[string(plaintext)] = ciphertext // Store the ciphertext in the cache.
	cache.Unlock()                          // Release the write lock.

	return ciphertext, nil
}

// AdvancedDecrypt decrypts a ciphertext and caches the result.
func AdvancedDecrypt(ciphertext []byte, key []byte) ([]byte, error) {
	cache.RLock()                                // Acquire a read lock.
	plaintext, ok := cache.m[string(ciphertext)] // Check if the plaintext is in the cache.
	cache.RUnlock()                              // Release the read lock.

	if ok {
		return plaintext, nil // If the plaintext is in the cache, return it.
	}

	// If the plaintext is not in the cache, decrypt the ciphertext.
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

	cache.Lock()                             // Acquire a write lock.
	cache.m[string(ciphertext)] = ciphertext // Store the plaintext in the cache.
	cache.Unlock()                           // Release the write lock.

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
