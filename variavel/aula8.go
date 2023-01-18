package main

import "fmt"

/* var x = "Hello World!" -> assim, fora do bloco abaixo, funciona como escopo*/ 

func main(){
	var x = "Hello World!"
	fmt.Println(x)

	/* outra maneira de escrever o mesmo código:
	
	var x string
	x = "Hello World!"
	fmt.Println(x)

	*/

	var y [10]int
	y[9] = 150
	fmt.Println(y)
	
}