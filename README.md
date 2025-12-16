# Go Rest

## Api endpoints

    http://localhost:8080/api
    http://localhost:8080/api/old
    http://localhost:8080/api/string
    http://localhost:8080/api/int
    http://localhost:8080/api/float
    http://localhost:8080/api/boolean
    http://localhost:8080/api/user?id=1
    http://localhost:8080/api/users

## Init

    go mod init go-rest
    
    go run src/restServer.go
    
    go run src/controller/homeController.go

## Build

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
