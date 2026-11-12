
package main

import (
	"fmt"
	"io"
	"os"

	"password-cracker/utils/crack"
)

func main() {

	targetHash := "6a85dfd77d9cb35770c9dc6728d73d3f"
	

	// Create verbose output file
	file, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Display output on terminal and save to file
	output := io.MultiWriter(os.Stdout, file)

	fmt.Fprintln(output, "Starting MD5 Password Cracking...")
	fmt.Fprintln(output, "Target Hash:", targetHash)
	fmt.Fprintln(output, "--------------------------------")

	password, err := crack.CrackMD5(
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
