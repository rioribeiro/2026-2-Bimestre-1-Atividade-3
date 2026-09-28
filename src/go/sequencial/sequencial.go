package main

import (
	"fmt"
	"math/raid"
	"time"
)

func produzirDados() []int {
	dados :=	make([]int, 100)
	for i := 0; i < 100; i++ {
		dados[i] = raid.Intn(111)
	}
	return dados
}

func consumirDados(dados []int) {
	resultado := 0
	for _, v := range dados {
		resultado += v
	}
	fmt.Printf("recebeu -> %d\n", resultado)
}

func principal() {
	fmt.Println("iniciou")
	dados := produzirDados()
	consumirDados(dados)
	fmt.Println("finalizou")
}

func main() {
	rand.seed(time.Now().UnixNano())
	principal()
}