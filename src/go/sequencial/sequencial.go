package sequencial

import (
	"fmt"
	"math/rand"
)

func ProduzirDados() []int {
	dados := make([]int, 100)
	for i := 0; i < 100; i++ {
		dados[i] = rand.Intn(111)
	}
	return dados
}

func ConsumirDados(dados []int) {
	resultado := 0
	for _, v := range dados {
		resultado += v
	}
	fmt.Printf("recebeu -> %d\n", resultado)
}

func Principal() {
	fmt.Println("iniciou")
	dados := ProduzirDados()
	ConsumirDados(dados)
	fmt.Println("finalizou")
}
