package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"strconv"
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
			n1, _ := strconv.Atoi(number1)
			n2, _ := strconv.Atoi(number2)			

		}

	}
	fmt.Println("Primeiro número:", number1)
	fmt.Println("Operação:", operation)
	fmt.Println("Segundo número:", number2)

}
func destiny(char rune) {
	if char == '+' {
		resultado := Soma(n1, n2)
	} else if char == '-' {
		resultado := Substracao(n1, n2)
	} else if char == '/' {
		resultado := Divisao(n1, n2)
	} else if char == '*' {
		resultado := Multiplicacao(n1, n2)
	}
}
