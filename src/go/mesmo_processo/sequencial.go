package main

import (
	"fmt"
	"math/rand"
	"time"
)

// Produz os dados e retorna o slice de inteiros
func produzirDados() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dados := make([]int, 100)
	for i := 0; i < 100; i++ {
		dados[i] = r.Intn(111)
	}
	return dados
}

// Recebe a lista por parâmetro e calcula a soma
func consumirDados(dados []int) {
	resultado := 0
	for _, num := range dados {
		resultado += num
	}
	fmt.Printf("Recebeu -> %d\n", resultado)
}

func principal() {
	fmt.Println("Iniciou")

	// Transfere o retorno de produzirDados() diretamente para consumirDados()
	dados := produzirDados()
	consumirDados(dados)

	fmt.Println("Finalizou")
}

func main() {
	principal()
}
