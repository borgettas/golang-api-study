package health

import (
	"time"
)

// Service define a lógica de negócio para o módulo de saúde (health).
type Service struct {
	startTime time.Time
}

// NewService cria uma nova instância do serviço de saúde,
// armazenando a hora de início da aplicação.
func NewService() *Service {
	return &Service{
		startTime: time.Now(),
	}
}

// GetUptime calcula e retorna o tempo de atividade da aplicação.
func (s *Service) GetUptime() time.Duration {
	return time.Since(s.startTime)
}
