package messages

import "log"

// Service define a lógica de negócio para o módulo de mensagens.
// Por enquanto, é uma estrutura simples sem campos.
// No futuro, poderia ter um campo de conexão com o banco de dados.
type Service struct {}

// NewService cria e retorna uma nova instância do serviço de mensagens.
func NewService() *Service {
	return &Service{}
}

// PrintMessage é a função de lógica de negócio.
// Ela recebe a string e a imprime no log.
func (s *Service) PrintMessage(content string) {
	log.Printf("String recebida: '%s'", content)
}