package sequencial

import (
	"fmt"
	"math/rand"
	"time"
)

// ProduzirDados (com P maiúsculo) é visível para outros pacotes
func ProduzirDados() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dados := make([]int, 100)
	for i := 0; i < 100; i++ {
		dados[i] = r.Intn(111)
	}
	return dados
}

// ConsumirDados recebe os dados e calcula a soma
func ConsumirDados(dados []int) {
	soma := 0
	for _, v := range dados {
		soma += v
	}
	fmt.Printf("recebeu -> %d\n", soma)
}

// Principal executa a lógica completa de produzir e consumir
func Principal() {
	fmt.Println("iniciou")
	dados := ProduzirDados()
	ConsumirDados(dados)
	fmt.Println("finalizou")
}