package parser

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RichardtJustke/calcGo/operacoes"
)

func InputUser() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Digite sua conta (ex: 2+2): ")

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
			continue
		}

		if readingSecondNumber {
			number2 += string(char)
		} else {
			number1 += string(char)
		}
	}

	number1 = strings.TrimSpace(number1)
	number2 = strings.TrimSpace(number2)

	if number1 == "" || number2 == "" || operation == "" {
		fmt.Println("Entrada inválida. Formato esperado: <num><op><num>, ex: 12+34")
		return
	}

	n1, err1 := strconv.Atoi(number1)
	n2, err2 := strconv.Atoi(number2)
	if err1 != nil || err2 != nil {
		fmt.Println("Não foi possível converter os números para inteiros.")
		return
	}

	var resultado int
	switch operation {
	case "+":
		resultado = operacoes.Soma(n1, n2)
	case "-":
		resultado = operacoes.Substracao(n1, n2)
	case "*":
		resultado = operacoes.Multiplicacao(n1, n2)
	case "/":
		if n2 == 0 {
			fmt.Println("Erro: divisão por zero")
			return
		}
		resultado = operacoes.Divisao(n1, n2)
	default:
		fmt.Println("Operador não suportado:", operation)
		return
	}

	fmt.Printf("Resultado: %d\n", resultado)
}
