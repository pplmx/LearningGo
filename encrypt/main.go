package main

import (
	"fmt"
	"github.com/google/uuid"
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

	// Encrypt the plaintextFiles concurrently
	err := lib.EncryptFiles("data", key)
	if err != nil {
		fmt.Println("Error during encryption: ", err)
		return
	}

	fmt.Println("Files encrypted successfully")
}

func decryptFilesDemo() {
	key := []byte("myverystrongpasswordo32bitlength") // replace with your key

	// Decrypt the ciphertextFiles concurrently
	err := lib.DecryptFiles("data", key)
	if err != nil {
		fmt.Println("Error during decryption: ", err)
		return
	}

	fmt.Println("Files decrypted successfully")
}

func encryptFilesCli() {
	binaryName := filepath.Base(os.Args[0])
	usage := fmt.Sprintf(`
Usage: %s <Command>

Commands:
	encrypt <file/directory>          The encrypted files will be stored in a directory named encrypted_<uuid>
	decrypt <file/directory>          The decrypted files will be stored in a directory named decrypted_<uuid>

Examples:
	%s encrypt data
	%s decrypt encrypted_12345678
	`, binaryName, binaryName, binaryName)

	if len(os.Args) < 3 {
		fmt.Println(usage)
		os.Exit(1)
	}

	op := os.Args[1]
	if op != "encrypt" && op != "decrypt" {
		fmt.Println(usage)
		os.Exit(1)
	}
	path := os.Args[2]

	// Generate a short UUID
	uuidStr := uuid.New().String()[:8]

	// Determine the destination directory based on the operation
	destDir := "encrypted_" + uuidStr
	if op == "decrypt" {
		destDir = "decrypted_" + uuidStr
	}

	// Create the destination directory if it doesn't exist
	if _, err := os.Stat(destDir); os.IsNotExist(err) {
		err := os.Mkdir(destDir, 0755)
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}
	}

	// Copy the files or directories to the destination directory
	err := lib.CopyFiles(path, destDir)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

	// Encrypt or decrypt the files in the destination directory
	if op == "encrypt" {
		err = lib.EncryptFiles(destDir, []byte("myverystrongpasswordo32bitlength"))
	} else {
		err = lib.DecryptFiles(destDir, []byte("myverystrongpasswordo32bitlength"))
	}

	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
