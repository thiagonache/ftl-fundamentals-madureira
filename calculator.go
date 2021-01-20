// Package calculator provides a library for simple calculations in Go.
package calculator

import (
	"errors"
	"math"
	"regexp"
	"strconv"
)

// Add takes two numbers and returns the result of adding them together.
func Add(a float64, b float64, c ...float64) float64 {
	var retorno float64 = a + b

	for _, item := range c {
		retorno += item
	}
	return retorno
}

// Subtract takes two numbers and returns the result of subtracting the second
// from the first.
func Subtract(a float64, b float64, c ...float64) float64 {
	var retorno float64 = a - b
	// for i := 0; i < len(c); i++ {
	// 	retorno -= c[i]
	// }
	for _, item := range c {
		retorno -= item
	}
	return retorno
}

// Multiply takes two numbers and returns the result
func Multiply(a float64, b float64, c ...float64) float64 {
	var retorno float64 = a * b
	// for i := 0; i < len(c); i++ {
	// 	retorno *= c[i]
	// }
	for _, item := range c {
		retorno *= item
	}
	return retorno
}

// Divide takes tow numbers and return the result
func Divide(a float64, b float64, c ...float64) (float64, error) {
	// return 2, errors.New("Anta não divide por 0 porra....")
	retorno := a
	if b == 0 {
		return 0, errors.New("Anta não divide por 0 porra....")
	}
	retorno /= b
	// for i := 0; i < len(c); i++ {
	// 	if c[i] == 0 {
	// 		return 0, errors.New("Anta não divide por 0 porra....")
	// 	}
	// 	retorno /= c[i]
	// }

	for _, item := range c {

		if item == 0 {
			return 0, errors.New("Anta não divide por 0 porra....")
		}
		retorno /= item
	}

	return retorno, nil
}

// Sqrt from a positive number.
func Sqrt(a float64) (float64, error) {
	if a < 0 {
		return 0, errors.New("Não Existe raiz quadrada de número negativo")
	}
	return math.Sqrt(a), nil
}

// ParseString receive a string and return the math cont
func ParseString(parse string) (float64, error) {
	reg := regexp.MustCompile(`([0-9]+(\.[0-9]+)?) *([\+\*-\/]) *([0-9]+(\.[0-9]+)?)`)
	val := reg.FindStringSubmatch(parse)
	num1, _ := strconv.ParseFloat(val[1], 64)
	num2, _ := strconv.ParseFloat(val[4], 64)
	if val[3] == "*" {
		return Multiply(num1, num2), nil
	} else if val[3] == "+" {
		return Add(num1, num2), nil
	} else if val[3] == "-" {
		return Subtract(num1, num2), nil
	} else if val[3] == "/" {
		return Divide(num1, num2)
	}
	return 0, errors.New("Inválido")
}

//func main() {
//	teste := " 3.3 / 1.3"
//	teste1, err := ParseString(teste)
//	fmt.Printf("%#v\n", teste1)
//	fmt.Printf("%#v\n", err)
//}
