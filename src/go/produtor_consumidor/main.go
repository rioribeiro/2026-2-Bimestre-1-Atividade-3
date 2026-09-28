package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

var (
	dados []int
	wg    sync.WaitGroup
)

func produzirDados() {
	defer wg.Done()
	fmt.Println("# produzir - iniciado")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	novosDados := make([]int, 100)
	for i := 0; i < 100; i++ {
		novosDados[i] = r.Intn(111)
	}
	dados = novosDados

	fmt.Printf("# produzir %v\n", dados)
	fmt.Println("# produzir - terminado")
}

func consumirDados() {
	defer wg.Done()
	fmt.Println("### consumir - iniciado")
	fmt.Printf("### dados -> %v\n", dados)

	soma := 0
	for _, v := range dados {
		soma += v
	}

	fmt.Printf("### resultado -> %d\n", soma)
	fmt.Println("### consumir - terminado")
}

func main() {
	fmt.Println("iniciou")

	wg.Add(2)
	go produzirDados()
	go consumirDados()

	wg.Wait()
	fmt.Println("finalizou")
}