package messages

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/go-sql-driver/mysql"
)

// Service gerencia a lógica de negócio do módulo de mensagens, incluindo a conexão com o banco de dados.
type Service struct {
	db *sql.DB
}

// NewService cria e retorna uma nova instância do serviço, estabelecendo a conexão com o banco de dados.
func NewService(user, password, host, port, name string) (*Service, error) {
	dataSourceName := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s", user, password, host, port, name)

	db, err := sql.Open("mysql", dataSourceName)
	if err != nil {
		return nil, fmt.Errorf("erro ao conectar ao banco de dados: %w", err)
	}

	// Testa a conexão
	if err = db.Ping(); err != nil {
		return nil, fmt.Errorf("erro ao fazer ping no banco de dados: %w", err)
	}

	log.Println("Conexão com o banco de dados MySQL estabelecida com sucesso!")
	return &Service{db: db}, nil
}

// SaveMessage salva uma string no banco de dados.
func (s *Service) SaveMessage(content string) error {
	insertStmt, err := s.db.Prepare("INSERT INTO messages (content) VALUES (?)")
	if err != nil {
		return fmt.Errorf("erro ao preparar a declaração SQL: %w", err)
	}
	defer insertStmt.Close()

	_, err = insertStmt.Exec(content)
	if err != nil {
		return fmt.Errorf("erro ao executar a declaração: %w", err)
	}

	return nil
}

// CloseDB fecha a conexão com o banco de dados.
func (s *Service) CloseDB() {
	if s.db != nil {
		s.db.Close()
	}
}

// DB retorna a conexão com o banco de dados.
func (s *Service) DB() *sql.DB {
	return s.db
}

