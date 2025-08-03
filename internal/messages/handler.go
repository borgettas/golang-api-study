package messages

import (
	"encoding/json"
	"log"
	"net/http"
)

// Handler define a estrutura para o manipulador HTTP do módulo de mensagens.
type Handler struct {
	service *Service
}

// NewHandler cria e retorna uma nova instância do handler.
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// Message representa a estrutura do corpo da requisição JSON.
type Message struct {
	Content string `json:"content"`
}

// Handle lida com as requisições HTTP e delega a lógica para o serviço.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Requisição JSON inválida", http.StatusBadRequest)
		return
	}

	// Chama o serviço para salvar a mensagem no banco de dados
	if err := h.service.SaveMessage(msg.Content); err != nil {
		http.Error(w, "Erro ao salvar a string", http.StatusInternalServerError)
		log.Printf("Erro ao salvar mensagem: %v", err)
		return
	}

	log.Printf("String salva com sucesso: '%s'", msg.Content)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"message": "String salva com sucesso no banco de dados!",
	}
	json.NewEncoder(w).Encode(response)
}

