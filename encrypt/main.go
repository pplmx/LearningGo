package main

import (
	"fmt"
	"github.com/pplmx/LearningGo/encrypt/lib"
	"os"
)

func main() {
	demo()
	fmt.Println("=====================================")
	advancedDemo()
	fmt.Println("=====================================")
	concurrentDemo()
	fmt.Println("=====================================")
	encryptFilesDemo()
	fmt.Println("=====================================")
	decryptFilesDemo()
}

func demo() {
	key := []byte("myverystrongpasswordo32bitlength") // replace with your key
	plaintext := []byte("Hello, World!")

	fmt.Println("Original text: ", string(plaintext))

	// Encrypt the plaintext
	ciphertext, err := lib.Encrypt(plaintext, key)
	if err != nil {
		fmt.Println("Error during encryption: ", err)
		return
	}

	fmt.Println("Encrypted text: ", ciphertext)

	// Decrypt the ciphertext
	decryptedText, err := lib.Decrypt(ciphertext, key)
	if err != nil {
		fmt.Println("Error during decryption: ", err)
		return
	}

	fmt.Println("Decrypted text: ", string(decryptedText))
}

func advancedDemo() {
	key := []byte("myverystrongpasswordo32bitlength") // replace with your key
	plaintext := []byte("Hello, World!")

	fmt.Println("Original text: ", string(plaintext))

	// Encrypt the plaintext
	ciphertext, err := lib.AdvancedEncrypt(plaintext, key)
	if err != nil {
		fmt.Println("Error during encryption: ", err)
		return
	}

	fmt.Println("Encrypted text: ", ciphertext)

	// Decrypt the ciphertext
	decryptedText, err := lib.AdvancedDecrypt(ciphertext, key)
	if err != nil {
		fmt.Println("Error during decryption: ", err)
		return
	}

	fmt.Println("Decrypted text: ", string(decryptedText))
}

func concurrentDemo() {
	key := []byte("myverystrongpasswordo32bitlength") // replace with your key
	plaintexts := [][]byte{[]byte("Hello World"), []byte("Goodbye World"), []byte("Hello Again")}

	fmt.Println("Original texts: ", plaintexts)
	for _, plaintext := range plaintexts {
		fmt.Println("Original text: ", string(plaintext))
	}

	// Encrypt the plaintexts concurrently
	ciphertexts, err := lib.ConcurrentEncrypt(plaintexts, key)
	if err != nil {
		fmt.Println("Error during encryption: ", err)
		return
	}

	fmt.Println("Encrypted texts: ", ciphertexts)

	// Decrypt the ciphertexts concurrently
	decryptedTexts, err := lib.ConcurrentDecrypt(ciphertexts, key)
	if err != nil {
		fmt.Println("Error during decryption: ", err)
		return
	}

	fmt.Printf("Decrypted texts: %v\n", decryptedTexts)

	for _, decryptedText := range decryptedTexts {
		fmt.Println("Decrypted text: ", string(decryptedText))
	}
}

func encryptFilesDemo() {
	key := []byte("myverystrongpasswordo32bitlength") // replace with your key
	// Define the file names
	fileNames := []string{"data/hi.txt", "data/hello.yaml", "data/hey.toml"}

	// Create a slice to store the file pointers
	plaintextFiles := make([]*os.File, len(fileNames))

	// Open each file and append it to the slice
	for i, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()

		plaintextFiles[i] = file
	}

	// Encrypt the plaintextFiles concurrently
	err := lib.EncryptFiles(plaintextFiles, key)
	if err != nil {
		fmt.Println("Error during encryption: ", err)
		return
	}

	fmt.Println("Files encrypted successfully")
}

func decryptFilesDemo() {
	key := []byte("myverystrongpasswordo32bitlength") // replace with your key
	// Define the file names
	fileNames := []string{"data/hi.txt", "data/hello.yaml", "data/hey.toml"}

	// Create a slice to store the file pointers
	ciphertextFiles := make([]*os.File, len(fileNames))

	// Open each file and append it to the slice
	for i, fileName := range fileNames {
		file, err := os.Open(fileName)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			return
		}
		defer file.Close()

		ciphertextFiles[i] = file
	}

	// Decrypt the ciphertextFiles concurrently
	err := lib.DecryptFiles(ciphertextFiles, key)
	if err != nil {
		fmt.Println("Error during decryption: ", err)
		return
	}

	fmt.Println("Files decrypted successfully")
}
