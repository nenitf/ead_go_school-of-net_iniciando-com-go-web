package main

import (
	"fmt"
	"net/http"
)

func main() {

	// rota "/"
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "hello")
	})

	fmt.Println(http.ListenAndServe(":8080", nil)) //Escutando porta 8080
}
