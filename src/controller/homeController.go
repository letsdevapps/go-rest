package main

import (
    "fmt"
    "io/ioutil"
    "log"
    "net/http"
)

func main() {
    //resp, err := http.Get("https://localhost:8080/api")
    resp, err := http.Get("http://localhost:8080/api")
    if err != nil {
        log.Fatal(err)
    }
    defer resp.Body.Close()

    body, err := ioutil.ReadAll(resp.Body)
    if err != nil {
        log.Fatal(err)
    }

    fmt.Println("Resposta:", string(body))
}

