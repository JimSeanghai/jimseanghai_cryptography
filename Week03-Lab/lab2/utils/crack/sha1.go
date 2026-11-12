
package crack

import (
	"bufio"
	"crypto/sha1"
	"encoding/hex"
	"fmt"
	"io"
	"os"
	"strings"
)

func CrackSHA1(filename string, targetHash string, output io.Writer) (string, error) {

	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	count := 0

	for scanner.Scan() {

		password := strings.TrimSpace(scanner.Text())

		hash := sha1.Sum([]byte(password))
		hashString := hex.EncodeToString(hash[:])

		count++

		fmt.Fprintln(output, "Trying:", password, "=>", hashString)

		if hashString == strings.ToLower(targetHash) {

			fmt.Fprintln(output, "\nPassword Found!")
			fmt.Fprintln(output, "Password:", password)
			fmt.Fprintln(output, "Attempts:", count)

			return password, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Fprintln(output, "\nPassword not found!")

	return "", nil
}
