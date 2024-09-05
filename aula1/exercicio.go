package main

import "fmt"
  
const ebulicao = 373.0

func main() {
	var tempK float64 = ebulicao
	var tempC float64 =tempK - 273
	fmt.Println("a escala da ebulição da água em °K =", tempK)
	fmt.Println("a escala da ebulição da água em °C =", tempC)
}

