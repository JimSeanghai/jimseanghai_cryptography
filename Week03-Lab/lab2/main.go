
package main

import (
	"fmt"
	"io"
	"os"

	"password-cracker/utils/crack"
)

func main() {

	targetHash := "aa1c7d931cf140bb35a5a16adeb83a551649c3b9"

	// Create verbose output file
	file, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Display output on terminal and save to file
	output := io.MultiWriter(os.Stdout, file)

	fmt.Fprintln(output, "Starting SHA1 Password Cracking...")
	fmt.Fprintln(output, "Target Hash:", targetHash)
	fmt.Fprintln(output, "--------------------------------")

	password, err := crack.CrackSHA1(
		"wordlists/nord_vpn.txt",
		targetHash,
		output,
	)

	if err != nil {
		fmt.Fprintln(output, "Error:", err)
		return
	}

	if password != "" {
		fmt.Fprintln(output, "\nCracking completed successfully!")
	} else {
		fmt.Fprintln(output, "\nCracking completed. Password not found.")
	}
}
