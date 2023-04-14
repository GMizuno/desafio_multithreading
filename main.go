package main

import (
	"github.com/GMizuno/desafio_multithreading/request"
)

func main() {
	Address1, _ := request.ApiCdn("24346-030")

	println(Address1)

	Address2, _ := request.ApiViaCep("24346-030")

	println(Address2)
}
