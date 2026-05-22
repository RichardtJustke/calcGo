package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite sua conta: ")

	conta, _ := reader.ReadString('\n')

	conta = strings.TrimSpace(conta)

	number1 := ""
	number2 := ""
	operation := ""

	readingSecondNumber := false
	for _, char := range conta {

		if char == '+' || char == '-' || char == '*' || char == '/' {

			operation = string(char)

			readingSecondNumber = true

		} else {

			if readingSecondNumber {

				number2 += string(char)

			} else {

				number1 += string(char)

			}

		}

	}
	fmt.Println("Primeiro número:", number1)
	fmt.Println("Operação:", operation)
	fmt.Println("Segundo número:", number2)

}
func destiny(char rune) {
	if char == '+' {
		fmt.Println("é uma soma")
	} else if char == '-' {
		fmt.Println("é uma subtração")
	} else if char == '/' {
		fmt.Println("é uma divisão")
	} else if char == '*' {
		fmt.Println("é uma multiplicação")
	}
}
