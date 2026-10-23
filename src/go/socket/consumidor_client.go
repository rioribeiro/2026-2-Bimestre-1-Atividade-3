packpage main

import(
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	// O nome do host 'produtor-host' será resolvido pelo DNS do Docker host := os Getenv ("PRODUTOR-HOST")
	if host == "" {
		host = "produtor-host"
	}
	porta := "5000"
	endereco := fmt.Sprintf("%s:%s", host, porta)

	fmt.Printf("### [Cliente/Consumidor] Tentando conectar ao produtor em %s\n", endereco)

	var conn net.conn
	var err error

	// Tenta reconectar caso o servidor ainda esteja inicializado 
	for i := 0; i < 10; i++ {
		conn, err = net.Dial("tcp", endereco)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}
	if err != nil {
		fmt.Printf("Erro no consumidor ao conectar: %v\n", err)
		return
	}
	defer conn.Close()

	// Lê a resposta enviada pelo servidor
	leitor := bufio.NewReader(conn)
	resposta, err := leitor.ReadString('\n')
	if err != nil {
		fmt.Printf("Erro ao ler dados do socket: %v\n", err)
		return
	}

	fmt.Printf("### [Cliente/Consumidor] Dados recebidos via Socket -> %s", resposta)
	// Trata a String "[1 2 3 ...]" de volta para números inteiros e calcula a soma
	limpo := strings.TrimSpace(resposta)
	limpo := strings.TrimPrefix(limpo, "[")
	limpo := strings.TrimSuffix(limpo, "]")

	partes := strings.Fields(limpo)
	resultado := 0

	for _, p := range partes {
		val, err := strconv.Atoi(p)
		if err == nil {
			resultado += val
		}
	}

	fmt.Printf("### [Cliente/Consumidor] Resultado da Soma -> %d\n", resultado)
}