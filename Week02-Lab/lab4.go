
package main

import "fmt"

func add(a, b float64) float64 {
	return a+b
}

func sub(a, b float64) float64 {
	return a - b
}

func mul( a, b float64) float64 {
	return a * b
}

func div( a, b float64) (float64, bool) {
	if b == 0 {
		return 0, false
	}
	return a/b, true
}

func mod(a, b int) (int, bool) {
	if b == 0 {
		return 0, false
	}
	return a % b, true
}

func main() {
	for {
		fmt.Println("\n===== Mini Calculator =====")
		fmt.Println("1) Add")
		fmt.Println("2) Sub")
		fmt.Println("3) Mul")
		fmt.Println("4) Div")
		fmt.Println("5) Mod")
		fmt.Println("6) Exit")

		var choice int
		fmt.Print("Choose: ")
		fmt.Scan( &choice)

		if choice==6 {
			fmt.Println(" Goodbye! ")
			break
		}

		if choice < 1 || choice> 6 {
			fmt.Println("Error: invalid choice")
			continue
		}

		if choice ==5 {
			var a,b int
			fmt.Print(" Enter a: ")
			fmt.Scan(&a)
			fmt.Print("Enter b: ")
			fmt.Scan( &b)

			result, ok := mod(a, b)
			if !ok {
				fmt.Println(" Error : division by zero")
			} else {
				fmt.Println("Result:", result )
			}
			continue
		}

		var a,b float64
		fmt.Print("Enter a: ")
		fmt.Scan(&a)
		fmt.Print(" Enter b: ")
		fmt.Scan( &b)

		switch choice {
		case 1:
			fmt.Println("Result:" , add(a, b))
		case 2:
			fmt.Println("Result:", sub(a, b))
		case 3:
			fmt.Println("Result: ", mul(a, b))
		case 4:
			result, ok :=div( a, b)
			if !ok {
				fmt.Println(" Error: division by zero")
			} else {
				fmt.Println("Result:", result)
			}
		}
	}
}
