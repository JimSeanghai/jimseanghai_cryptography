package main

import "fmt"

func main(){

	var a int = 2
	var b int = 2

	fmt.Printf("Initialized Vaule a = %d and b = %d \n", a,b)

	fmt.Printf( "Perform Operation = : a = b => %d \n", a)

	a+=b
	fmt.Printf("Perform Operation += : a += b => %d \n", a )

	a-=b
	fmt.Printf("Perform Operation -= : a -= b => %d \n", a)
	
	a*=b
	fmt.Printf("Perform Operation *= : a *= b => %d \n", a)

	a/=b
	fmt.Printf("Perform Operation /= : a /= b => %d \n", a)
	
	if b!=0{
		a%=b
		fmt.Printf("Perform Operation %= : a %= b => %d \n", a)
	}else {
		fmt.Printf("Modulo can not devide by Zero")
	}

}