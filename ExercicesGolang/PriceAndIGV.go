package main

import (
	"fmt"
	"strconv"
)

func PriceAndIGV() {
	var n1 string
	fmt.Println("Leer numeros")
	fmt.Println("Escribe el valor de venta del producto:")
	fmt.Scanln(&n1)
	numbConvert1, err1 := strconv.Atoi(n1)
	if err1 != nil {
		fmt.Println("Error during conversion")
		return
	}

	calPri := float64(numbConvert1)
	calIgv := calPri * 0.18

	fmt.Printf("El valor de venta del producto es: $%v \n", calPri)
	fmt.Printf("El IGV es: $%v \n", calIgv)
	fmt.Printf("El precio de venta es: $%v \n", calPri+calIgv)
}
