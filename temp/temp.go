package main

import "fmt"

// declaração da variável CONST do ponto de ebulição da água em F
const ebulicaoF = 212.0

func main(){
	tempF := ebulicaoF // variável do valor da temperatura em graus F
	tempC := (tempF - 32)*5/9 // conversão de F para C

	// apareça na tela o resultado
	fmt.Println("Ponto de ebulição da água em °F é: ", tempF)
  fmt.Println("Ponto de ebulição da água em °C é: ", tempC)

	// para usar o printf (f de format):
	// fmt.Printf("Ponto de ebulição da água em °F é: %g  e Ponto de ebulição da água em °C é: %g ", tempF, tempC)

	// espera-se que apareça F = 212 e C = 100

}