FROM golang:latest

WORKDIR /app

COPY . .

##Check if go.mod exists, if not it should create
##RUN go mod init go-rest
##RUN [ -f go.mod ] || go mod init go-rest
RUN if [ ! -f go.mod ]; then go mod init go-rest; fi

RUN mkdir -p bin

RUN go build -o bin/go-rest ./src

EXPOSE 8080

CMD ["./bin/go-rest"]
