package main

import (
	"fmt"
	"log"
	"strconv"
)

func declaredVar() {

	// declaracion de constantes
	const pi float64 = 3.14
	const pi2 = 3.16
	fmt.Println("Hello-word")
	fmt.Println("pi: ", pi, "pi2: ", pi2)

	// declaracion de variables enteras
	base := 12 //estamos diciendo que se creara la variable y le asignaremos un valor
	var altura int = 14
	var area int = 2
	fmt.Println(base, altura, area)

	// Zero value
	var a int
	var b float64
	var c string
	var d bool
	fmt.Println(a, b, c, d)

	// Area cuadrado
	const baseC = 10
	areaC := baseC * baseC
	fmt.Println("area de cuadrado: ", areaC)

	x := 16
	y := 8

	//suma
	result := x + y
	fmt.Println("Suma: ", result)

	//resta
	result = x - y
	fmt.Println("Resta: ", result)

	//multiplicacion
	result = x * y
	fmt.Println("Multiplicacion: ", result)

	//division
	result = x / y
	fmt.Println("division: ", result)

	for i := 0; i < 3; i++ {
		x++
		y--
	}

	fmt.Println("X: ", x, "Y: ", y)

	// Printf
	nombre := "Platzi"
	cursos := 500
	// Con valores seguros
	fmt.Printf("%s tiene más de %d cursos\n", nombre, cursos)
	// Con valores inseguros
	fmt.Printf("%v tiene más de %v cursos\n", nombre, cursos)

	// Sprintf
	message := fmt.Sprintf("%v tiene más de %v cursos\n", nombre, cursos)
	fmt.Println(message)

	// Tipo de datos:
	fmt.Printf("helloMessage: %T\n", nombre)
	fmt.Printf("cursos: %T\n", cursos)
}

func printHello(message string) string {
	value_return := fmt.Sprintf("Hello %s", message)
	return value_return
}

func returnTwo(a, b int) (z, y int) {
	z = a * 3
	y = b * 4
	return z, y
}

func moreFunctions() {

	value := printHello("no se xd")
	fmt.Println(value)

	valueZ, valueY := returnTwo(4, 5)
	var result int = valueZ + valueY
	printear := fmt.Sprintf("%d mas %d es igual a %d", valueZ, valueY, result)

	primerValue, _ := returnTwo(2, 2)

	fmt.Print(printear, " segunda func: ", primerValue, "\n")
}

func forTypes() {

	//for condicional
	for i := 0; i < 10; i++ {
		fmt.Println(i)
	}

	// for while
	counter := 100
	for counter > 0 {
		fmt.Println("for while ", counter)
		counter--
	}

	//for forever

	counterForever := 0
	for {
		fmt.Println(counterForever)
		counterForever++
	}
}

func condicionales(textInt string) {
	//convertir texto a numero con if else
	value, e := strconv.Atoi(textInt)
	if e != nil {
		log.Fatal(e)
	}
	fmt.Println(value)
}

func isPar(n int) bool {
	if (n % 2) == 0 {
		return true
	} else {
		return false
	}
}

func contadorVocales(palabra string) (int, int, int, int, int) {
	conta := 0
	conte := 0
	conti := 0
	conto := 0
	contu := 0
	for _, valor := range palabra {
		variable := string(valor)
		switch variable {
		case "a":
			conta++
		case "e":
			conte++
		case "i":
			conti++
		case "o":
			conto++
		case "u":
			contu++
		}
	}
	return conta, conte, conti, conto, contu
}

func withDefer() {
	defer fmt.Println("defer") // se ejecuta al finalizar el programa
	fmt.Println("normal")
}

func withBreakAndContinue() {
	defer fmt.Println("defer")
	for i := 0; i < 10; i++ {
		if i == 5 {
			continue
		}
		if i == 7 {
			break
		}
		fmt.Println(i)
	}
}
func main() {
	withBreakAndContinue()
}
