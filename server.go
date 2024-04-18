package main

import (
	"encoding/json"
	"net/http"
)

func StartServer() {
	http.HandleFunc("/produtos", produtosHandler) // coloca o que ta provendo em plural normalmente, e depois o handler
	// 8080 porta para iniciar, para fazer o servidor funcionar
	http.ListenAndServe(":8080", nil)
}

func produtosHandler(w http.ResponseWriter, r *http.Request) {
	produtos := catalogo()

	// estamos pedindo para o json fazer o encode do produto
	json.NewEncoder(w).Encode(produtos) // encoder pega um objeto e traduz para json
}
