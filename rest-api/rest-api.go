package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type PayLoad struct {
	Stuff Data `json:"stuff"`
}

type Data struct {
	Fruit   Fruits     `json:"fruit"`
	Veggies Vegetables `json:"veggies"`
}

type Fruits map[string]int
type Vegetables map[string]int

func serveRest(w http.ResponseWriter, r *http.Request) {
	response, err := getJsonResponse()
	if err != nil {
		panic(err)
	}

	fmt.Fprint(w, string(response))
}

func main() {
	http.HandleFunc("/", serveRest)
	http.ListenAndServe("localhost:1234", nil)
}

func getJsonResponse() ([]byte, error) {
	fruits := make(map[string]int)
	fruits["Apples"] = 25
	fruits["Oranges"] = 11

	vegetables := make(map[string]int)
	vegetables["Carrots"] = 21
	vegetables["Peppers"] = 0

	d := Data{fruits, vegetables}
	p := PayLoad{d}
	return json.MarshalIndent(p, "", " ")
}
