package health

import (
	"encoding/json"
	"net/http"
)

// Handler define a estrutura para o manipulador HTTP do módulo de saúde (health).
type Handler struct {
	service *Service
}

// NewHandler cria e retorna uma nova instância do handler.
func NewHandler(s *Service) *Handler {
	return &Handler{service: s}
}

// Handle lida com as requisições HTTP para a rota de saúde.
func (h *Handler) Handle(w http.ResponseWriter, r *http.Request) {
	// Obtém o tempo de atividade e o status do banco de dados
	uptime := h.service.GetUptime()
	dbStatus := h.service.CheckDBHealth()

	response := map[string]interface{}{
		"success": true,
		"status":  "UP",
		"uptime":  uptime.String(),
		"dependencies": map[string]string{
			"database": dbStatus,
		},
	}

	w.WriteHeader(http.StatusOK)
	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(response)
}
