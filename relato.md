# Relatório de Implementação de Comunicação entre Tarefas em Go

## 📌 Introdução

Este relato faz parte do processo avaliativo da disciplina de **Sistemas Operacionais** no curso superior de **Tecnologia em Análise e Desenvolvimento de Sistemas**, ofertado pela **Diretoria Acadêmica de Gestão e Tecnologia da Informação** do **Campus Natal-Central** do **Instituto Federal de Educação, Ciência e Tecnologia do Rio Grande do Norte (**IFRN**)**.

O objetivo principal deste trabalho é relatar e demonstrar as implementações de **comunicação entre tarefas** na linguagem **Go**.

## 👥 Equipe de Desenvolvimento

- **Rio Ribeiro**
- **Wheverton Filho**

---

## ⚙️ Comunicação entre Tarefas em Go

### 💡 Informações Gerais

A comunicação entre tarefas (***IPC** - Inter-Process Communication* ou comunicação inter-goroutines) tem como objetivo permitir que diferentes fluxos de execução troquem dados, sincronizem suas operações e cooperem para resolver um problema em comum. Em arquiteturas baseadas no padrão **Produtor-Consumidor**, essa abordagem possibilita a separação de responsabilidades entre a geração de dados e seu processamento.

### 🐳 Uso e Configuração do Docker

O uso do **Docker** neste trabalho garante a padronização do ambiente de execução entre os membros da equipe e na avaliação do professor. O contêiner isola o ambiente da linguagem Go (`golang:1.22-alpine`), garantindo que limitações de sistema operacional, variáveis de ambiente ou diferenças de versão não interferiram na criação e no agendamento de Goroutines e processos.

---

## 🛠️ Relato de Desenvolvimento e Implementações

### 🧵 1. Comunicação entre Tarefas no Mesmo Processo (Linhas de Execução)

A comunicação entre tarefas dentro do mesmo processo em Go ocorre por meio de **Goroutines** (threads leves gerenciadas pelo runtime da linguagem), que compartilham o mesmo espaço de endereçamento de memória **RAM**. Nos códigos implementados, essa integração foi demonstrada em três cenários graduais:

- **Chamada Direta de Fluxo (`exemplo.go`):** Execução sequencial de funções sem transferência de estado.
- **Troca por Parâmetros e Retornos (`sequencial.go`):** A comunicação ocorre pela passagem de fatias de memória (*slices*) através do fluxo único de controle.
- **Memória Compartilhada (`produtor_consumidor.go`):** Múltiplas Goroutines concorrentes acessam e alteram uma variável global comum (`dados []int`), com sincronização controlada via `sync.WaitGroup`.

#### 💻 Códigos Implementados:

1. **Chamada Direta de Fluxo (`exemplo.go`)**
```go
package main

import *fmt*

func produzirDados() {
	fmt.Println(*Dados produzidos*)
}

func main() {
	produzirDados()
}
```

2. **Comunicação por Parâmetro e Retorno (`sequencial.go`)**
```go
package main

import (
	*fmt*
	*math/rand*
	*time*
)

func produzirDados() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	dados := make([]int, **100**)
	for i := 0; i < **100**; i++ {
		dados[i] = r.Intn(**111**)
	}
	return dados
}

func consumirDados(dados []int) {
	resultado := 0
	for _, num := range dados {
		resultado += num
	}
	fmt.Printf(*Recebeu -> %d\n*, resultado)
}

func principal() {
	fmt.Println(*Iniciou*)
	dados := produzirDados()
	consumirDados(dados)
	fmt.Println(*Finalizou*)
}

func main() {
	principal()
}
```

3. **Comunicação Concorrente via Memória Compartilhada (`produtor_consumidor.go`)**
```go
package main

import (
	*fmt*
	*math/rand*
	*sync*
	*time*
)

var dados []int

func produzirDados(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(*# Produzir - iniciado*)

	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	novosDados := make([]int, **100**)
	for i := 0; i < **100**; i++ {
		novosDados[i] = r.Intn(**111**)
	}

	dados = novosDados

	fmt.Printf(*# produzir %v\n*, dados)
	fmt.Println(*# produzir - terminado*)
}

func consumirDados(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(*### Consumir - iniciado*)
	fmt.Printf(*### Dados -> %v\n*, dados)

	resultado := 0
	for _, num := range dados {
		resultado += num
	}

	fmt.Printf(*### Resultado -> %d\n*, resultado)
	fmt.Println(*### Consumir - terminado*)
}

func principal() {
	fmt.Println(*iniciou*)

	var wg sync.WaitGroup

	wg.Add(1)
	go produzirDados(&wg)
	wg.Wait()

	wg.Add(1)
	go consumirDados(&wg)
	wg.Wait()

	fmt.Println(*Finalizou*)
}

func main() {
	principal()
}
```

#### 📥 Execução do Código via Docker:

```
sudo docker build -t app-go-mesmo-processo .
sudo docker run --rm app-go-mesmo-processo
```

**Saída no terminal do contêiner:** 
```
iniciou 
# Produzir - iniciado 
# produzir [34 93 93 68 40 37 95 91 ...] 
# produzir - terminado 
### Consumir - iniciado 
### Dados -> [34 93 93 68 40 37 95 91 ...]
### Resultado -> 5786 
### Consumir - terminado 
Finalizou 
```

#### ⚠️ Desafios Encontrados e Soluções:

1. **Condição de Corrida (*Race Condition*):**
    * **Problema:** Ao disparar Goroutines simultâneas sem sincronização, a Goroutine consumidora tentava ler a variável global `dados` antes que o produtor tivesse terminado de preenchê-la.
    * **Solução:** Uso do mecanismo de barreira de sincronização `sync.WaitGroup` com a instrução `wg.Wait()`, garantindo a ordem correta de execução.

---

### 💻 2. Comunicação entre Tarefas em Processos Diferentes no Mesmo Computador (Pipe)

Nesta etapa, o padrão é adaptado para executar a produção e o consumo em **processos completamente isolados** no Sistema Operacional.
- **Isolamento de Memória:** Cada processo possui seu próprio espaço de endereçamento protegido pelo SO.
- **Mecanismo via Pipe (`|`):** A comunicação é feita redirecionando a saída padrão (*stdout*) do processo produtor para a entrada padrão (*stdin*) do processo consumidor.

#### 💻 Códigos Implementados:

1. **Produtor (`produtor.go`)**
```go
package main

import *fmt*

func main() {
	fmt.Println(*Dados produzidos e enviados via Pipe de SO*)
}
```

2. **Consumidor (`consumidor.go`)**
```go
package main

import (
	*bufio*
	*fmt*
	*os*
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	for scanner.Scan() {
		fmt.Println(*Entrada recebida no Consumidor ->*, scanner.Text())
	}
}
```

#### 📥 Execução do Código via Docker:

```bash 
sudo docker build -f Dockerfile.produtor -t app-produtor . 
sudo docker build -f Dockerfile.consumidor -t app-consumidor .
sudo docker run --rm app-produtor | sudo docker run -i --rm app-consumidor
```

**Saída no terminal do contêiner:** 
``` 
Entrada recebida no Consumidor -> Dados produzidos e enviados via Pipe de SO 
```

#### ⚠️ Desafios Encontrados e Soluções:

1. **Perda da Entrada Padrão no Contêiner Consumidor:**
    * **Problema:** O contêiner Docker do consumidor por padrão não mantinha a entrada interativa aberta para ler dados do pipe.
    * **Solução:** Uso obrigatório da flag `-i` no comando `docker run -i`, mantendo o fluxo `stdin` aberto para receber o fluxo de dados do produtor.

---

### 🌐 3. Comunicação entre Tarefas em Computadores/Contêineres Diferentes (Sockets TCP)

Para evoluir a arquitetura para um ambiente distribuído, o código foi dividido em dois programas independentes que se comunicam através da pilha de protocolos **TCP**/IP:
- **`produtor_server.go` (Servidor Socket):** Atua como o nó gerador de dados, abrindo uma porta de rede **TCP** (`**5000**`) e aguardando conexões.
- **`consumidor_client.go` (Cliente Socket):** Conecta-se ao hostname do produtor via socket, consome a stream de bytes transmitida via rede, realiza o parsing do vetor e calcula a soma.
- **Orquestração com Docker Compose:** Gerencia a rede virtual (`bridge`) conectando os dois contêineres e resolvendo os nomes via **DNS** interno do Docker.

#### 💻 Código Implementado:

1. **Contêiner A (`produtor_server.go`)**
```go
package main

import (
	*fmt*
	*math/rand*
	*net*
	*time*
)

func produzirDados() []int {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	novosDados := make([]int, **100**)
	for i := 0; i < **100**; i++ {
		novosDados[i] = r.Intn(**111**)
	}
	return novosDados
}

func main() {
	porta := *:**5000***
	fmt.Println("# [Servidor/Produtor] Aguardando conexões na porta **5000**...*)

	listener, err := net.Listen(*tcp", porta)
	if err != nil {
		return
	}
	defer listener.Close()

	for {
		conn, err := listener.Accept()
		if err != nil {
			continue
		}

		fmt.Printf(*# [Servidor/Produtor] Consumidor conectado: %s\n*, conn.RemoteAddr().String())

		dados := produzirDados()
		mensagem := fmt.Sprintf(*%v\n*, dados)

		conn.Write([]byte(mensagem))
		fmt.Printf("# [Servidor/Produtor] Dados enviados via Socket: %v\n", dados)
		conn.Close()
	}
}
```

2. **Contêiner B (`consumidor_client.go`)**
```go
package main

import (
	*bufio*
	*fmt*
	*net*
	*os*
	*strconv*
	*strings*
	*time*
)

func main() {
	host := os.Getenv(*PRODUTOR_HOST*)
	if host == "* {
		host = *produtor-host*
	}
	porta := *5000*
	endereco := fmt.Sprintf(*%s:%s*, host, porta)

	fmt.Printf(*### [Cliente/Consumidor] Tentando conectar ao produtor em %s\n", endereco)

	var conn net.Conn
	var err error

	for i := 0; i < 10; i++ {
		conn, err = net.Dial(*tcp*, endereco)
		if err == nil {
			break
		}
		time.Sleep(1 * time.Second)
	}

	if err != nil {
		return
	}
	defer conn.Close()

	leitor := bufio.NewReader(conn)
	resposta, _ := leitor.ReadString('\n')

	fmt.Printf("### [Cliente/Consumidor] Dados recebidos via Socket -> %s", resposta)

	limpo := strings.TrimSpace(resposta)
	limpo = strings.TrimPrefix(limpo, *[*)
	limpo = strings.TrimSuffix(limpo, *]*)

	partes := strings.Fields(limpo)
	resultado := 0

	for _, p := range partes {
		val, err := strconv.Atoi(p)
		if err == nil {
			resultado += val
		}
	}

	fmt.Printf(*### [Cliente/Consumidor] Resultado da Soma -> %d\n*, resultado)
}
```

3. **Orquestração (`docker-compose.yml`)**
```yaml
services:
    produtor-host:
    build: .
    container_name: conteiner_produtor
    command: go run produtor_server.go
    ports:
    - ***5000**:**5000***
    networks:
    - rede_ipc

    consumidor-host:
    build: .
    container_name: conteiner_consumidor
    command: go run consumidor_client.go
    environment:
    - PRODUTOR_HOST=produtor-host
    depends_on:
    - produtor-host
    networks:
    - rede_ipc

networks:
    rede_ipc:
    driver: bridge
```

#### 📥 Execução do Código via Docker Compose:

```sudo docker-compose up --build ```

**Saída do terminal do contêiner:** 
``` 
conteiner_produtor  | # [Servidor/Produtor] Aguardando conexões na porta 5000...
conteiner_consumidor  | ### [Cliente/Consumidor] Tentando conectar ao produtor em produtor-host:5000
conteiner_produtor    | # [Servidor/Produtor] Consumidor conectado: 172.19.0.3:37034
conteiner_produtor    | # [Servidor/Produtor] Dados enviados via Socket: [15 100 49 108 2 32 35 93 54 ...]
conteiner_consumidor  | ### [Cliente/Consumidor] Dados recebidos via Socket -> [15 100 49 108 2 32 35 93 54 ...]
conteiner_consumidor  | ### [Cliente/Consumidor] Resultado da Soma -> 5233
conteiner_consumidor exited with code 0
```

#### ⚠️ Desafios Encontrados e Soluções:

1. **Ordem de Inicialização dos Contêineres:**
    * **Problema:** O consumidor tentava conectar via Socket antes do servidor aceitar conexões **TCP**.
    * **Solução:** Implementação de um laço de tentativa de reconexão (`retry loop` com `time.Sleep`) no cliente Go e diretiva `depends_on` no Compose.
2. **Serialização e Parsing do Fluxo de Bytes:**
    * **Problema:** A rede transmite os dados como fluxo de bytes (`[]byte`).
    * **Solução:** Formatação da fatia em string no servidor e reconstrução dos valores numéricos com `strings.Fields` e `strconv.Atoi` no cliente.

---

## 🏁 Considerações Finais

A realização deste trabalho permitiu vivenciar, de forma prática e incremental, a evolução das abordagens de Comunicação entre Tarefas em Go:

1. **Memória Compartilhada no Mesmo Processo:** Alta performance, porém exige atenção rigorosa ao controle de concorrência com Goroutines e travas para evitar condições de corrida.
2. **Processos Isolados no Mesmo Computador:** Garante o isolamento completo de memória pelo SO, utilizando Pipes para redirecionamento do fluxo de dados.
3. **Rede Distribuída com Sockets:** Abstrai a dependência de memória local e habilita a arquitetura em microsserviços via Docker Compose.

