
package main

import "fmt"

func main () {

	var a , b  , c float64
	fmt.Println("enter a")
	fmt.Scanln(& a)

	fmt.Println("enter b")
	fmt.Scanln(& b)

	c = a * b
	fmt.Printf("value of a & b : %f \n" ,c )
}