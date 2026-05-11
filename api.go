package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Cliente struct {
	ID    string `json:"id"`
	Nome  string `json:"nome"`
	Email string `json:"email"`
}

var clientes []Cliente

// Handlers (As funções que processam as requisições)

func obterClientes(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(clientes)
}

func obterCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r) // Pega os parâmetros da URL
	for _, item := range clientes {
		if item.ID == params["id"] {
			json.NewEncoder(w).Encode(item)
			return
		}
	}
	w.WriteHeader(http.StatusNotFound)
}

func criarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var cliente Cliente
	_ = json.NewDecoder(r.Body).Decode(&cliente)
	clientes = append(clientes, cliente)
	json.NewEncoder(w).Encode(cliente)
}

func deletarCliente(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	params := mux.Vars(r)
	for index, item := range clientes {
		if item.ID == params["id"] {
			clientes = append(clientes[:index], clientes[index+1:]...)
			break
		}
	}
	w.WriteHeader(http.StatusNoContent)
}

func main() {
	// Inicializa o roteador do Gorilla Mux
	router := mux.NewRouter()

	// Mock de dados inicial
	clientes = append(clientes, Cliente{ID: "1", Nome: "Fulano de Tal", Email: "fulano@email.com"})

	// Definição das Rotas e Verbos HTTP
	router.HandleFunc("/clientes", obterClientes).Methods("GET")
	router.HandleFunc("/clientes/{id}", obterCliente).Methods("GET")
	router.HandleFunc("/clientes", criarCliente).Methods("POST")
	router.HandleFunc("/clientes/{id}", deletarCliente).Methods("DELETE")

	fmt.Println("Servidor rodando na porta 8080...")
	log.Fatal(http.ListenAndServe(":8080", router))
}
