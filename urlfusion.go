package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Lê o parâmetro da linha de comando
	args := os.Args[1:]
	if len(args) == 0 {
		fmt.Println("Por favor, forneça um parâmetro.")
		return
	}
	parametro := args[0]

	// Lê os valores da lista via pipeline
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		valor := scanner.Text()
		resultado := valor + parametro
		fmt.Println(resultado)
	}

	if err := scanner.Err(); err != nil {
		fmt.Fprintln(os.Stderr, "Erro ao ler a entrada padrão:", err)
	}
}
