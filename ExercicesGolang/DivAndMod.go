package main

import (
	"fmt"
	"math"
	"strconv"
)

func main() {
	var n1, n2 string
	fmt.Println("Leer numeros")
	fmt.Println("Escribe el numero uno:")
	fmt.Scanln(&n1)
	fmt.Println("Escribe el numero dos:")
	fmt.Scanln(&n2)
	numbConvert1, err1 := strconv.Atoi(n1)
	if err1 != nil {
		fmt.Println("Error during conversion")
		return
	}
	numbConvert2, err2 := strconv.Atoi(n2)
	if err2 != nil {
		fmt.Println("Error during conversion")
		return
	}

	calDiv := float64(numbConvert1) / float64(numbConvert2)

	fmt.Println("La division es:", calDiv)

	calMod := math.Mod(float64(numbConvert1), float64(numbConvert2))

	fmt.Println("La residuo es:", calMod)

	calMod2 := numbConvert1 % numbConvert2
	fmt.Println("La residuo con % es:", calMod2)
}
