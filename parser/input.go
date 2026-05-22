package parse

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func inputUser() {
	reader := bufio.NewReader(os.Stdin)
	fmt.Printf("Digite sua conta:")
	conta, _ := reader.ReadString('\n')
	conta = strings.TrimSpace(conta)
	for _, char := range conta {
		fmt.Println(string(char))
		destiny(char)
	}

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
