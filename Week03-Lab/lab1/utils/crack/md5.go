
package crack

import (
	"bufio"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"os"
	"io"
	"strings"
)

func CrackMD5(filename string, targetHash string, output io.Writer) (string, error) {
	file, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)

	count := 0

	for scanner.Scan() {

		password := strings.TrimSpace(scanner.Text())

		hash := md5.Sum([]byte(password))
		hashString := hex.EncodeToString(hash[:])

		count++

		//fmt.Println("Trying:", password, "=>", hashString)
		fmt.Fprintln(output, "Trying: ", password, "=>", hashString)

		if hashString == strings.ToLower(targetHash) {

			// fmt.Println("\nPassword Found!")
			// fmt.Println("Password:", password)
			// fmt.Println("Attempts:", count)
			fmt.Fprintln(output,"\nPassword Found!")
			fmt.Fprintln(output,"Password:", password)
			fmt.Fprintln(output,"Attempts:", count)

			return password, nil
		}
	}

	if err := scanner.Err(); err != nil {
		return "", err
	}

	fmt.Fprintln(output,"\nPassword not found!")
	return "", nil
}
