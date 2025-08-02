package health

import (
	"encoding/json"
	"net/http"
)

// Handler define a estrutura para o manipulador HTTP do módulo de saúde (health).
// Ele tem uma dependência do serviço de saúde para obter o tempo de atividade.
type Handler struct {
	service *Service
}

// NewHandler cria e retorna uma nova instância do handler.
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// Handle é o método principal que lida com as requisições HTTP para a rota de saúde.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Obtém o tempo de atividade do serviço
	uptime := h.service.GetUptime()

	// Cria a resposta JSON com o tempo de atividade em um formato de string
	response := map[string]interface{}{
		"success":    true,
		"status":     "UP",
		"components": nil,
		"uptime":     uptime.String(),
	}

	// Envia a resposta
	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
