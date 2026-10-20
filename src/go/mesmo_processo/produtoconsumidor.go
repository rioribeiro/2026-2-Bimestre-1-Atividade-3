package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// Memória compartilhada entre Goroutines dentro do mesmo processo
var dados []int

func produzirDados(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("# Produzir - iniciado")

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	novosDados := make([]int, 100)
	for i := 0; i < 100; i++ {
		novosDados[i] = r.Intn(111)
	}

	dados = novosDados // Atualiza a variável global comum

	fmt.Printf("# produzir %v\n", dados)
	fmt.Println("# produzir - terminado")
}

func consumirDados(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("### Consumir - iniciado")
	fmt.Printf("### Dados -> %v\n", dados)

	resultado := 0
	for _, num := range dados {
		resultado += num
	}

	fmt.Printf("### Resultado -> %d\n", resultado)
	fmt.Println("### Consumir - terminado")
}

func principal() {
	fmt.Println("iniciou")

	var wg sync.WaitGroup

	// Executa a Goroutine produtora e aguarda o término (equivalente ao thread.join() do Java)
	wg.Add(1)
	go produzirDados(&wg)
	wg.Wait()

	// Executa a Goroutine consumidora e aguarda o término (equivalente ao thread.join() do Java)
	wg.Add(1)
	go consumirDados(&wg)
	wg.Wait()

	fmt.Println("Finalizou")
}

func main() {
	principal()
}
