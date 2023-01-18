package main

import "fmt"

const ebulicaoK = 373.0

func main(){

	tempK := ebulicaoK
	tempC := tempK - 273.0

	fmt.Printf("A temperatura de ebulição da água em °K = %g (%T) e a temperatura de ebulição da água em °C = %g (%T)", tempK, tempK, tempC, tempC)

}