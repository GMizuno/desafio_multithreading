package main

import (
	"fmt"
	"github.com/GMizuno/desafio_multithreading/request"
	"time"
)

func main() {
	c1 := make(chan string)
	c2 := make(chan string)
	cep := "24346030"

	go func(cep string) {
		address, _ := request.ApiCdn(cep)
		c1 <- address
	}(cep)

	go func(cep string) {
		address, _ := request.ApiViaCep(cep)
		c2 <- address
	}(cep)

	select {
	case address := <-c1:
		fmt.Println("Usando api Cdn, temos:\n", address)
	case address := <-c2:
		fmt.Println("Usando api ViaCep, temos:\n", address)
	case <-time.After(time.Second * 1):
		println("timeout")
	}
}
