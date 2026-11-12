
package main

import (
	"fmt"
	"io"
	"os"

	"password-cracker/utils/crack"
)

func main() {

	targetHash := "485f5c36c6f8474f53a3b0e361369ee3e32ee0de2f368b87b847dd23cb284b316bb0f026ada27df76c31ae8fa8696708d14b4d8fa352dbd8a31991b90ca5dd38"

	// Create verbose output file
	file, err := os.Create("verbose.txt")
	if err != nil {
		fmt.Println("Error creating file:", err)
		return
	}
	defer file.Close()

	// Display output on terminal and save to file
	output := io.MultiWriter(os.Stdout, file)

	fmt.Fprintln(output, "Starting SHA512 Password Cracking...")
	fmt.Fprintln(output, "Target Hash:", targetHash)
	fmt.Fprintln(output, "--------------------------------")

	password, err := crack.CrackSHA512(
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
