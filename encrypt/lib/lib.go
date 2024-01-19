package lib

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"crypto/sha256"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"
	"sync"
)

var cache = struct {
	sync.RWMutex
	m map[string][]byte
}{m: make(map[string][]byte)}

func EncryptFiles(path string, key []byte) error {
	var wg sync.WaitGroup

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()

				file, err := os.Open(path)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer file.Close()

				err = encryptFile(file, key)
				if err != nil {
					fmt.Println(err)
					return
				}
			}(path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func encryptFile(file *os.File, key []byte) error {
	// Get the file info.
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Read the file into a buffer.
	plaintext := make([]byte, info.Size())
	_, err = file.Read(plaintext)
	if err != nil {
		return err
	}

	// Encrypt the plaintext.
	ciphertext, err := AdvancedEncrypt(plaintext, key)
	if err != nil {
		return err
	}

	// Write the ciphertext to the file.
	err = os.WriteFile(file.Name(), ciphertext, info.Mode())
	if err != nil {
		return err
	}

	return nil
}

func DecryptFiles(path string, key []byte) error {
	var wg sync.WaitGroup

	err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		if !info.IsDir() {
			wg.Add(1)
			go func(path string) {
				defer wg.Done()

				file, err := os.Open(path)
				if err != nil {
					fmt.Println(err)
					return
				}
				defer file.Close()

				err = decryptFile(file, key)
				if err != nil {
					fmt.Println(err)
					return
				}
			}(path)
		}
		return nil
	})

	if err != nil {
		return err
	}

	wg.Wait()
	return nil
}

func decryptFile(file *os.File, key []byte) error {
	// Get the file info.
	info, err := file.Stat()
	if err != nil {
		return err
	}

	// Read the file into a buffer.
	ciphertext := make([]byte, info.Size())
	_, err = file.Read(ciphertext)
	if err != nil {
		return err
	}

	// Decrypt the ciphertext.
	plaintext, err := AdvancedDecrypt(ciphertext, key)
	if err != nil {
		return err
	}

	// Write the plaintext to the file.
	err = os.WriteFile(file.Name(), plaintext, info.Mode())
	if err != nil {
		return err
	}

	return nil
}

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

// XOR a simple encrypt/decrypt algorithm
// a ^ b ^ b = a
// secret, such as 0xff
func XOR(raw []byte, secret byte) []byte {
	for i := range raw {
		raw[i] ^= secret
	}
	return raw
}

func Hash(data []byte) [32]byte {
	return sha256.Sum256(data)
}
