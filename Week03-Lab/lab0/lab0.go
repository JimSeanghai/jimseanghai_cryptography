package main

import (
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"crypto/sha512"
	"encoding/hex"
	"fmt"
	"golang.org/x/crypto/sha3"
)

func verify(txt1, txt2 string)  {
	//md5
	hashA := md5.Sum([]byte(txt1))
	hashB := md5.Sum([]byte(txt2))
	checkHash("MD5", hashA[:], hashB[:])


	// sha1
	hashC := sha1.Sum([]byte(txt1))
	hashD := sha1.Sum([]byte(txt2))
	checkHash("SHA1", hashC[:], hashD[:])
	
	//sha256
	hashE := sha256.Sum256([]byte(txt1))
	hashF := sha256.Sum256([]byte(txt2))
	checkHash("SHA256", hashE[:], hashF[:])

	// sha512
	hashG := sha512.Sum512([]byte(txt1))
	hashH := sha512.Sum512([]byte(txt2))
	checkHash("SHA512", hashG[:], hashH[:])
	
	// sha3-256
	hashI := sha3.Sum256([]byte(txt1))
	hashJ := sha3.Sum256([]byte(txt2))
	checkHash("SHA3-256", hashI[:], hashJ[:] )
}


func checkHash(name string, outputA, outputB []byte) {

	fmt.Println("\nHash (" + name + "):")
	fmt.Println("Output A =", hex.EncodeToString(outputA))
	fmt.Println("Output B =", hex.EncodeToString(outputB))

	if hex.EncodeToString(outputA) == hex.EncodeToString(outputB) {
		fmt.Println("=> Match!")
	} else {
		fmt.Println("=> No Match!")
	}
}

func main() {

	var Input1, Input2 string

	fmt.Println("======== Name + Hashing Program ========")

	fmt.Print("Please input value 1: ")
	fmt.Scanln(&Input1)

	fmt.Print("Please input value 2: ")
	fmt.Scanln(&Input2)

	verify(Input1, Input2)
}
