package health

import (
	"database/sql"
	"time"
)

// Service define a lógica de negócio para o módulo de saúde (health).
type Service struct {
	startTime time.Time
	db        *sql.DB
}

// NewService cria uma nova instância do serviço de saúde,
// recebendo a conexão com o banco de dados.
func NewService(db *sql.DB) *Service {
	return &Service{
		startTime: time.Now(),
		db:        db,
	}
}

// GetUptime calcula e retorna o tempo de atividade da aplicação.
func (s *Service) GetUptime() time.Duration {
	return time.Since(s.startTime)
}

// CheckDBHealth verifica o status da conexão com o banco de dados.
func (s *Service) CheckDBHealth() string {
	if s.db == nil {
		return "DOWN"
	}
	if err := s.db.Ping(); err != nil {
		return "DOWN"
	}
	return "UP"
}

