# Relatório sobre implementação de comunicação entre tarefas em Go

## Introdução

Este relato faz parte do processo avaliativo da disciplina de sistemas operacionas no curso superior em análise e desenvolvimento de sistemas, ofertado na Diretoria acadêmica de gestão e tecnologia da informação no campus natal-central do instituto federal de educação, ciência e tecnologia do rio grande do norte.

Tem como objetivo principal relatar as implementações de comunicação entre tarefas na linguagem Go.

O grupo de trabalho foi formado por Rio Ribeiro, João Victor, Wheverton Filho.

## Comunicação entre tarefas em Go

### Informações gerais

A comunicação e a coordenação entre tarefas (sejam processos ou *threads/goroutines*) têm como objetivo permitir a troca de dados, a sincronização do fluxo de execução e a cooperação para resolver problemas complexos. Em Sistemas Operacionais, ela permite dividir uma aplicação grande em submódulos que executam concorrentemente, garantindo o compartilhamento controlado de recursos (como memória e arquivos) e evitando inconsistências causadas por acessos simultâneos não protegidos.

O uso do **Docker** garante a portabilidade e a reprodutibilidade do ambiente de execução entre todos os membros do grupo e o professor. Ele isola as dependências do sistema operacional hospedeiro, garantindo que o código rode exatamente na mesma versão do Go (1.22-alpine), eliminando inconsistências de ambiente.

```
# Dockerfile

FROM golang:1.22-alpine AS builder 
WORKDIR /app 
COPY go/go.mod ./
COPY go/sequencial/ ./sequencial/
COPY go/exemplo\_main/ ./exemplo\_main/
COPY go/sequencial\_app/ ./sequencial\_app/
COPY go/produtor\_consumidor/ ./produtor\_consumidor/
RUN go build -o /app/bin/exemplo\_main ./exemplo\_main 
RUN go build -o /app/bin/sequencial\_app ./sequencial\_app 
RUN go build -o /app/bin/produtor\_consumidor ./produtor\_consumidor 
FROM alpine:latest 
WORKDIR /app 
COPY --from=builder /app/bin /app/bin
CMD ["/app/bin/produtor\_consumidor"]

# docker-compose.yml

services: 
  exemplo-main: 
    build: . 
    command: /app/bin/exemplo\_main 
    
  sequencial: 
    build: . 
    command: /app/bin/sequencial\_app
  produtor-consumidor:
    build: .
    command: /app/bin/produtor\_consumidor
```

### Comunicação entre tarefas com linhas de execução no mesmo processo

Em Go, linhas de execução no mesmo processo são gerenciadas através de Goroutines (threads leves gerenciadas pelo runtime da linguagem). No arquivo src/go/produtor_consumidor/main.go, implementamos a comunicação entre o Produtor e o Consumidor através de uma variável global compartilhada (dados []int) e controlamos a conclusão das Goroutines via barreira de sincronização sync.WaitGroup.
```
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
```

A execução foi realizada via terminal dentro do diretório do módulo Go (src/go):
```go run ./produtor_consumidor```
Saída obtida no terminal:

```
Iniciou
### consumir - iniciado
### dados -> []
### resultado -> 0
### consumir - terminado
# produzir - iniciado
# produzir 
# produzir - terminado
finalizou
```

Problemas na execução e soluções:

- Problema: Ocorreu uma Condição de Corrida (Race Condition). Como não havia mecanismos de exclusão mútua na memória compartilhada dados, a Goroutine consumirDados() executou antes que produzirDados() preenchesse o vetor. O consumidor leu um vetor vazio ([]) e resultou em soma 0.
- Solução: Para evitar a condição de corrida no Go, o ideal é substituir a memória compartilhada global por Canais (chan []int) para transmissão síncrona de dados entre as Goroutines, ou aplicar um sync.Mutex ao redor do acesso à variável dados.

### Comunicação entre tarefas em processos diferentes no mesmo computador

FIXME
> texto explicando o código
> mostrar o código completo

FIXME
> explicar como foi executado
> mostrar as saídas do terminal
> mostrar as saídas do terminal

FIXME
> se houve problema na execução, enumerar os problemas e suas respectivas soluções

### Comunicação entre tarefas em processos diferentes em computadores diferentes

FIXME
> texto explicando o código
> mostrar o código completo

FIXME
> explicar como foi executado
> mostrar as saídas do terminal
> mostrar as saídas do terminal

FIXME
> se houve problema na execução, enumerar os problemas e suas respectivas soluções

## Considerações finais

FIXME
> conseguiu implementar tudo e executar?
> qual foi o aprendizado nesse trabalho?
> alguma recomendação para próximos alunos?
