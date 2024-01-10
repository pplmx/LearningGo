package main

import (
	"fmt"
	"github.com/pplmx/LearningGo/encrypt/lib"
	"os"
	"path/filepath"
)

func main() {
	//demo()
	//fmt.Println("=====================================")
	//advancedDemo()
	//fmt.Println("=====================================")
	//concurrentDemo()
	//fmt.Println("=====================================")
	//encryptFilesDemo()
	//fmt.Println("=====================================")
	//decryptFilesDemo()

	encryptFilesCli()
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

func encryptFilesCli() {
	if len(os.Args) < 3 {
		fmt.Println("Usage: security encrypt/decrypt <file/directory>")
		os.Exit(1)
	}

	op := os.Args[1]
	if op != "encrypt" && op != "decrypt" {
		fmt.Println("Usage: security encrypt/decrypt <file/directory>")
		os.Exit(1)
	}
	path := os.Args[2]

	// Check if the path is a directory
	pathInfo, err := os.Stat(path)
	if err != nil {
		fmt.Println("Error opening file/directory: ", err)
		os.Exit(1)
	}

	var files []*os.File
	defer func() {
		for _, file := range files {
			file.Close()
		}
	}()

	// If the path is a directory, encrypt all files in the directory
	if pathInfo.IsDir() {
		// using filepath.Walk to recursively walk through all files in the directory
		err := filepath.Walk(path, func(path string, info os.FileInfo, err error) error {
			if info.IsDir() {
				return nil // skip directories
			}

			file, err := os.Open(path)
			if err != nil {
				return err
			}

			files = append(files, file)

			return nil
		})
		if err != nil {
			fmt.Println("Error walking through directory: ", err)
			os.Exit(1)
		}
	} else {
		file, err := os.Open(path)
		if err != nil {
			fmt.Println("Error opening file: ", err)
			os.Exit(1)
		}

		files = append(files, file)
	}

	if op == "encrypt" {
		err = lib.EncryptFiles(files, []byte("myverystrongpasswordo32bitlength"))
	} else {
		err = lib.DecryptFiles(files, []byte("myverystrongpasswordo32bitlength"))
	}
}
