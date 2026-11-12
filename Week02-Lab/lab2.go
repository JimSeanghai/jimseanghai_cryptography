package main

import "fmt"



func main() {

	var a int = 5
	var b int = 5
	var check bool 

	if (a > 0 && b > 0){
		check = true
		fmt.Printf("%t \n", check)
	}else{
		check = false
		fmt.Printf("%t \n", check)
	}

	if (a < b || b > a){
		check = true
		fmt.Printf("%t \n", check)
	}else{
		check = false
		fmt.Printf("%t \n", check)
	}

	if (a != b){
		check = true
		fmt.Printf("%t \n", check)
	}else{
		check = false
		fmt.Printf("%t \n", check)
	}

	
}
