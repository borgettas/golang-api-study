package messages

import (
	"encoding/json"
	"net/http"
)

// Handler define a estrutura para o manipulador HTTP do módulo de mensagens.
// Ele tem uma dependência do serviço de mensagens para executar a lógica de negócio.
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

// Handle é o método principal que lida com as requisições HTTP.
// Sua única responsabilidade é tratar a requisição e a resposta,
// delegando a lógica de negócio para o serviço.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Verifica se o método HTTP é POST
	if r.Method != http.MethodPost {
		http.Error(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Decodifica a requisição JSON
	var msg Message
	if err := json.NewDecoder(r.Body).Decode(&msg); err != nil {
		http.Error(w, "Requisição JSON inválida", http.StatusBadRequest)
		return
	}

	// Chama a lógica de negócio no serviço de mensagens
	h.service.PrintMessage(msg.Content)

	// Envia a resposta de sucesso
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{"status": true, "mensagem": "Received"}
	json.NewEncoder(w).Encode(response)
}
