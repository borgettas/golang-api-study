package main

import (
	"fmt"
	"log"
	"net/http"
	
	// Importa os módulos internos da aplicação
	"golang-api-study/internal/health"
	"golang-api-study/internal/messages"
)

func main() {
	// Inicializa os serviços
	messageService := messages.NewService()
	healthService := health.NewService()

	// Inicializa os handlers (manipuladores HTTP)
	messageHandler := messages.NewHandler(messageService)
	healthHandler := health.NewHandler(healthService)

	// Define os endpoints da API e atribui os handlers
	http.HandleFunc("/putter", messageHandler.Handle)
	http.HandleFunc("/health", healthHandler.Handle)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

