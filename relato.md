# Relatório sobre implementação de comunicação entre tarefas em Go

## Introdução

Este relato faz parte do processo avaliativo da disciplina de sistemas operacionas no curso superior em análise e desenvolvimento de sistemas, ofertado na Diretoria acadêmica de gestão e tecnologia da informação no campus natal-central do instituto federal de educação, ciência e tecnologia do rio grande do norte.

Tem como objetivo principal relatar as implementações de comunicação entre tarefas na linguagem Go.

O grupo de trabalho foi formado por Rio Ribeiro, João Victor, Wheverton Filho.

## Comunicação entre tarefas em Go

### Informações gerais

> qual o objetivo de comunicação entre tarefas?
A comunicação e a coordenação entre tarefas (sejam processos ou *threads/goroutines*) têm como objetivo permitir a troca de dados, a sincronização do fluxo de execução e a cooperação para resolver problemas complexos. Em Sistemas Operacionais, ela permite dividir uma aplicação grande em submódulos que executam concorrentemente, garantindo o compartilhamento controlado de recursos (como memória e arquivos) e evitando inconsistências causadas por acessos simultâneos não protegidos.
> explicar porque usar docker nesse trabalho.
O uso do **Docker** garante a portabilidade e a reprodutibilidade do ambiente de execução entre todos os membros do grupo e o professor. Ele isola as dependências do sistema operacional hospedeiro, garantindo que o código rode exatamente na mesma versão do Go (1.22-alpine), eliminando inconsistências de ambiente.
> qual a configuração do docker?
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

FIXME
> texto explicando o código
> mostrar o código completo

FIXME
> explicar como foi executado
> mostrar as saídas do terminal
> mostrar as saídas do terminal

FIXME
> se houve problema na execução, enumerar os problemas e suas respectivas soluções

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
