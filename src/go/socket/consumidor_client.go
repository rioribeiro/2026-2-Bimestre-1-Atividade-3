package main

import (
	"bufio"
	"fmt"
	"net"
	"os"
	"strconv"
	"strings"
	"time"
)

func main() {
	host := os.Getenv("PRODUTOR_HOST")
	if host == "" {
		host = "produtor-host"
	}
	porta := "5000"
	endereco := fmt.Sprintf("%s:%s", host, porta)

	fmt.Printf("### [Cliente/Consumidor] Tentando conectar ao produtor em %s\n", endereco)

	var conn net.Conn
	var err error

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

	leitor := bufio.NewReader(conn)
	resposta, err := leitor.ReadString('\n')
	if err != nil {
		fmt.Printf("Erro ao ler dados do socket: %v\n", err)
		return
	}

	fmt.Printf("### [Cliente/Consumidor] Dados recebidos via Socket -> %s", resposta)
	limpo := strings.TrimSpace(resposta)
	limpo = strings.TrimPrefix(limpo, "[")
	limpo = strings.TrimSuffix(limpo, "]")

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
