package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"golang-api-study/internal/health"
	"golang-api-study/internal/messages"
)

func main() {
	// Acessa as variáveis de ambiente do Docker Compose
	dbUser := os.Getenv("DB_USER")
	dbPass := os.Getenv("DB_PASSWORD")
	dbHost := os.Getenv("DB_HOST")
	dbPort := os.Getenv("DB_PORT")
	dbName := os.Getenv("DB_NAME")

	// Inicializa os serviços
	messageService, err := messages.NewService(dbUser, dbPass, dbHost, dbPort, dbName)
	if err != nil {
		log.Fatalf("Erro ao inicializar o serviço de mensagens: %v", err)
	}
	defer messageService.CloseDB()

	healthService := health.NewService(messageService.DB())

	// Inicializa os handlers
	messageHandler := messages.NewHandler(messageService)
	healthHandler := health.NewHandler(healthService)

	// Define os endpoints da API
	http.HandleFunc("/putter", messageHandler.Handle)
	http.HandleFunc("/health", healthHandler.Handle)

	fmt.Println("Servidor iniciado em http://localhost:8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}