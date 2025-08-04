package messages

import (
	"database/sql"
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

// Message representa o formato de resposta para o endpoint GET.
// Campos que podem ser nulos no banco de dados são agora representados por tipos do pacote sql.
type MessageResponse struct {
    ID        int            `json:"id"`
    Name      sql.NullString `json:"name"`
    Email     sql.NullString `json:"email"`
    Phone     sql.NullInt64  `json:"phone"`
    CreatedAt string         `json:"created_at"`
}


// ErrorResponse representa o formato padrão para as respostas de erro da API.
type ErrorResponse struct {
	Success bool   `json:"success"`
	Message string `json:"message"`
}

// sendErrorResponse é uma função de ajuda para enviar respostas de erro padronizadas.
func sendErrorResponse(w http.ResponseWriter, message string, statusCode int) {
	w.WriteHeader(statusCode)
	w.Header().Set("Content-Type", "application/json")
	response := ErrorResponse{
		Success: false,
		Message: message,
	}
	json.NewEncoder(w).Encode(response)
}

// Handle lida com as requisições HTTP e delega a lógica para o serviço.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		sendErrorResponse(w, "Método não permitido", http.StatusMethodNotAllowed)
		return
	}

	// Usa um map para decodificar o JSON e lidar com a natureza dinâmica da função de serviço.
	var dynamicMsg map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&dynamicMsg); err != nil {
		sendErrorResponse(w, "Requisição JSON inválida: "+err.Error(), http.StatusBadRequest)
		return
	}

	// Chama o serviço para salvar os dados no banco de dados
	if err := h.service.SaveDynamicMessages(dynamicMsg); err != nil {
		sendErrorResponse(w, "Erro ao salvar a string no banco de dados: "+err.Error(), http.StatusInternalServerError)
		log.Printf("Erro ao salvar mensagem: %v", err)
		return
	}

	log.Printf("Dados salvos com sucesso: %v", dynamicMsg)

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	response := map[string]interface{}{
		"success": true,
		"message": "Dados salvos com sucesso no banco de dados!",
	}
	json.NewEncoder(w).Encode(response)
}

// GetMessagesHandler lida com as requisições GET para o endpoint de mensagens.
func (h *Handler) GetMessagesHandler(w http.ResponseWriter, r *http.Request) {
    if r.Method != http.MethodGet {
        sendErrorResponse(w, "Método não permitido", http.StatusMethodNotAllowed)
        return
    }

    messages, err := h.service.GetMessages()
    if err != nil {
        sendErrorResponse(w, "Erro ao buscar mensagens: "+err.Error(), http.StatusInternalServerError)
        return
    }

    w.WriteHeader(http.StatusOK)
    w.Header().Set("Content-Type", "application/json")
    json.NewEncoder(w).Encode(messages)
}
