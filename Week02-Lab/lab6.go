
package main

import "fmt"

func xorEncrypt(text string, key byte) string {
	result := ""

	for i := 0; i< len(text); i++ {
		result +=string(text[i] ^ key)
	}

	return result
}

func main() {
	var text string
	var key byte

	fmt.Print("Enter plaintext: ")
	fmt.Scanln(&text)

	fmt.Print("Enter key (one character): ")
	fmt.Scanf(" %c", &key)

	// Encrypt message
	encrypted := xorEncrypt(text, key)
	fmt.Println("Encrypted text:", []byte( encrypted ))

	//decrypt the message
	decrypted := xorEncrypt(encrypted, key)
	fmt.Println(" Decrypted text:", decrypted)
}
