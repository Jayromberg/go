package main

import (
	"fmt"
	"reflect"
)

func main() {
	name := "Jayro"
	age := 31
	version := 1.1
	fmt.Println("Olá sr.", name, "sua idade é", age)
	fmt.Println("Este programa está na versão", version)

	fmt.Println("O tipo da variável name é", reflect.TypeOf(name))
	fmt.Println("O tipo da variável age é", reflect.TypeOf(age))
	fmt.Println("O tipo da variável version é", reflect.TypeOf(version))
}
