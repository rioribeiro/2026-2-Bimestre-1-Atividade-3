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