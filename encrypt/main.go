package main

import (
	"fmt"
	"github.com/pplmx/LearningGo/encrypt/lib"
)

func main() {
	simpleDemo()
	fmt.Println("=====================================")
	concurrentDemo()
}

func simpleDemo() {
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
