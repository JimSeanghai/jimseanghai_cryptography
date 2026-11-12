
package main

import "fmt"

func myXor(a, b int) {
    fmt.Println("XOR:", a ^ b)
}

func myNOT(a, b int) {
    fmt.Println("NOT a:", ^a)
    fmt.Println("NOT b:", ^b)
}

func myOR(a, b int) {
    fmt.Println("OR:", a | b)
}

func myAND(a, b int) {
    fmt.Println("AND:", a & b)
}

func main() {
    var a int = 4
	var b int = 6

    myXor(a, b)
    myNOT(a, b)
    myOR(a, b)
    myAND(a, b)

    fmt.Println("Left shift a:", a << 1)
    fmt.Println("Right shift a:", a >> 1)
}
