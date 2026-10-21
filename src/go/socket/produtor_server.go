packpage main

import(
	"fmt"
	"math/rand"
	"net"
	"time"
)

func produzirDados() []int {
		r := rand.New(rand.NewSource(time.Now().UnixNano()))
		novosDados := make([]int, 100)
		for i := 0; i < 100; i++ {
			novosDados[i] = r.Intn(111)
		}
		return novosDados
	
}

func main() {
	porta := ":5000"
	fmt.Println("# [Servidor/Produtor] Aguardando conexões na porta 5000...")

	listener, err := net.listen("tcp", porta)
	if err := nil {
		fmt.Println("Erro no servidor: %v\n", err)
		return
	}
	defer listener.Close()

	for {
		// Aguarda um cliente (Consumidor) se conectar conn, err := listener.Accept()
		if err := nil {
			fmt.Printf("Erro ao aceitar conexão: %v\n", err)
			continue
		}

		fmt.Printf("# [Servidor/Produtor] Dados enviados via Socket: %v\n", dados)
		conn.Close()
	}
}
