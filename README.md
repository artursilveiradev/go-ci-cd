# Go CI/CD

[![Go CI/CD with Docker](https://github.com/artursilveiradev/go-ci-cd/actions/workflows/ci-cd.yml/badge.svg)](https://github.com/artursilveiradev/go-ci-cd/actions/workflows/ci-cd.yml)

Este projeto é um exemplo de utilização de uma GitHub Action que compila e executa os testes de uma aplicação Go e publica a imagem Docker no Docker Hub.

## Construindo e executando

Para construir a imagem Docker:

```bash
docker build -t artursilveiradev/go-ci-cd .
```

Para executar o contêiner:

```bash
docker run artursilveiradev/go-ci-cd
```
