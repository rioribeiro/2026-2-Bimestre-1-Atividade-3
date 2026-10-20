package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Cria um leitor para a entrada padrão (stdin)
	scanner := bufio.NewScanner(os.Stdin)

	// Lê linha por linha até o encerramento do fluxo
	for scanner.Scan() {
		fmt.Println("Entrada recebida no Consumidor ->", scanner.Text())
	}
}
