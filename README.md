# Sistema de Produção e Consumo Kafka em Go

Este é um projeto simples que demonstra a implementação de um sistema de produção e consumo usando o Apache Kafka e a linguagem de programação Go.

## Pré-requisitos

Certifique-se de ter os seguintes pré-requisitos instalados em seu ambiente de desenvolvimento:

- Go: https://golang.org/doc/install
- Docker: https://docs.docker.com/get-docker/
- Docker Compose: https://docs.docker.com/compose/install/

## Configuração do Ambiente

Certifique-se de que o ambiente do Docker esteja em execução. Você pode iniciar o Kafka e o Zookeeper usando o seguinte comando:

```bash
docker-compose up -d

cd producer
go run main.go

cd consumer
go run main.go


