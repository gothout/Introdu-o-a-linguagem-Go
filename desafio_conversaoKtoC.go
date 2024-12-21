// Um professor de ensino médio solicitou aos seus alunos que desenvolvessem um programa para converter o valor do ponto de ebulição da agua de Kelvin para graus Celsius.
package main

import "fmt"

func main(){
	//Valor em K
	var K float64 = 373.15
	fmt.Printf("O valor convertido de %.2f Kelvin para Celsius é: %.2f",K,convertKtoC(K))
}

// Função que converte K em C, retorna float64
func convertKtoC(k float64) float64{
	// Efetuar a conversão
	var C float64 = (k - 273)
	return C
}