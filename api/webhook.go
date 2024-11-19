package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	// Processar o webhook
	if r.Method != http.MethodPost {
		http.Error(w, "Método não suportado", http.StatusMethodNotAllowed)
		return
	}

	var payload map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&payload); err != nil {
		http.Error(w, "Erro ao decodificar JSON", http.StatusBadRequest)
		return
	}

	fmt.Fprintf(w, "Webhook recebido: %v", payload)
}

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":3000", nil) // Não será usado na Vercel, mas mantém compatibilidade local
}
