# calcGo

Uma pequena calculadora em Go — projeto de aprendizado.

## Sobre

Este repositório é um exercício simples onde aprendi a implementar uma calculadora de linha de comando em Go. O objetivo foi praticar parsing de entrada, organização em pacotes e chamadas de funções.

## Funcionalidades

- Suporta operações básicas: soma, subtração, multiplicação e divisão.
- Entrada simples no formato: `12+34` (sem espaços) ou com espaços (`12 + 34`).
- Tratamento de erros básicos (divisão por zero, conversão inválida).

## Como rodar

Abra um terminal na raiz do projeto e execute:

```bash
go run .
```

Digite uma expressão no prompt, por exemplo:

```
12+5
```

E o programa retornará o resultado.

## O que aprendi

- Criação e organização de pacotes em Go (`parser`, `operacoes`).
- Leitura de entrada do usuário com `bufio`.
- Conversão de strings para inteiros e validação.
- Como compilar e executar programas Go simples.

## Validação na conversão de strings

Ao converter a entrada do usuário para números usamos `strconv.Atoi`, que retorna um erro se a string não for um inteiro válido. Validar esse erro evita que entradas inválidas sejam interpretadas como `0` silenciosamente — assim o programa avisa o usuário e evita cálculos incorretos.

## Mudanças futuras

- Suportar números decimais (`float`).
- Tratar sinais negativos e parênteses.
- Repetir o prompt até o usuário fornecer entrada válida.
- Adicionar testes unitários para as operações.
- Criar interface simples (web ou TUI).
---

Feito com ❤️ como parte do meu aprendizado em Go.
