# Go Rest

![GitHub release](https://img.shields.io/github/v/release/letsdevapps/go-rest)
![GitHub last commit](https://img.shields.io/github/last-commit/letsdevapps/go-rest)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/letsdevapps/go-rest/build-ci.yml?label=status%20integration)
![GitHub Actions Workflow Status](https://img.shields.io/github/actions/workflow/status/letsdevapps/go-rest/build-cd.yml?label=status%20deployment)

![Go](https://img.shields.io/badge/go-1.22+-brightgreen)

![Docker](https://img.shields.io/badge/docker-enabled-blue)
![CI](https://img.shields.io/badge/ci-enabled-blue)
![CD](https://img.shields.io/badge/cd-enabled-blue)
![Kubernetes](https://img.shields.io/badge/kubernetes-enabled-blue)
![Status](https://img.shields.io/badge/status-active-success)

## Docker Build

    docker build -t go-rest .

## Docker Run

    docker run --rm -it -p 8080:8080 go-rest

## Simple Run

    go run src/main.go
    
    go run src/restServer.go
    
    go run src/controller/homeController.go

## Init Run

    go mod init go-rest

## Build Run

Criar a pasta de saida (bin, target, build ...)

    mkdir -p bin

Aqui voce compila 1 unica classe
    
    go build -o bin/go-rest src/main.go

Aqui voce compila o projeto todo

    go build -o bin/go-rest .

Ou recursivamente abaixo da pasta especificada

    go build -o bin/go-rest ./src

    ./bin/go-rest

Para um ambiente de produção, você pode considerar compilar com a variável de ambiente GOOS e GOARCH para gerar binários específicos para diferentes sistemas operacionais e arquiteturas

    GOOS=linux GOARCH=amd64 go build -o bin/go-rest-linux src/main.go

## Comandos úteis do Go:

Criar um módulo Go: go mod init <nome_do_modulo>

Rodar um arquivo Go: go run <arquivo.go>

Instalar dependências: go get <url_do_repositorio>

Compilar um arquivo Go: go build <arquivo.go>

Verificar dependências: go mod tidy

## API endpoints

    GET http://localhost:8080/
    GET /api
    GET /api/string
    GET /api/int
    GET /api/float
    GET /api/boolean
    GET /api/user?id=1
    GET /api/users

