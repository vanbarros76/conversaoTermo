package main

import "fmt"

var a int = 20
var b string = "Nome"

func main(){
	var c float64 = float64(a) // tipo(variavel)
	fmt.Printf("O valor de c é: %g (%T) e o valor de b é: %s (%T)", c, c, b, b)
}