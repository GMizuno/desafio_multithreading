package request

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"time"
)

type ResponseCdn struct {
	Code       string `json:"code"`
	State      string `json:"state"`
	City       string `json:"city"`
	District   string `json:"district"`
	Address    string `json:"address"`
	Status     int    `json:"status"`
	Ok         bool   `json:"ok"`
	StatusText string `json:"statusText"`
}

type ResponseViaCep struct {
	Cep         string `json:"cep"`
	Logradouro  string `json:"logradouro"`
	Complemento string `json:"complemento"`
	Bairro      string `json:"bairro"`
	Localidade  string `json:"localidade"`
	Uf          string `json:"uf"`
	Ibge        string `json:"ibge"`
	Gia         string `json:"gia"`
	Ddd         string `json:"ddd"`
	Siafi       string `json:"siafi"`
}

func requester(url string) ([]byte, error) {
	client := http.Client{Timeout: 10 * time.Second}

	resp, err := client.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return nil, err
	}
	return body, nil
}

func ApiCdn(cep string) (string, error) {

	url := "https://cdn.apicep.com/file/apicep/" + cep + ".json"
	body, err := requester(url)
	var c ResponseCdn
	err = json.Unmarshal(body, &c)
	if err != nil {
		return "", err
	}

	return c.Address, nil

}

func ApiViaCep(cep string) (string, error) {

	url := "http://viacep.com.br/ws/" + cep + "/json/"
	body, err := requester(url)
	if err != nil {
		panic(err)
	}

	var c ResponseViaCep
	err = json.Unmarshal(body, &c)
	if err != nil {
		return "", err
	}

	return c.Logradouro, nil
}
