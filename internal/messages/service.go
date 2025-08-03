package messages

import (
	"database/sql"
	"fmt"
	"log"
	"strings"

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
// func (s *Service) SaveMessage(content string) error {
// func (s *Service) SaveMessage(content string) error {
// 	insertStmt, err := s.db.Prepare("INSERT INTO messages (content) VALUES (?)")
// 	if err != nil {
// 		return fmt.Errorf("erro ao preparar a declaração SQL: %w", err)
// 	}
// 	defer insertStmt.Close()

// 	_, err = insertStmt.Exec(content)
// 	if err != nil {
// 		return fmt.Errorf("erro ao executar a declaração: %w", err)
// 	}

// 	return nil
// }


func (s *Service) SaveDynamicMessages(data map[string]interface{}) error {
	// Constroi listas de colunas e placeholders dinamicamente
    cols := make([]string, 0, len(data))
    values := make([]interface{}, 0, len(data))
    placeholders := make([]string, 0, len(data))

    for col, val := range data {
        cols = append(cols, col)
        values = append(values, val)
        placeholders = append(placeholders, "?")
    }

    // Constroi a query SQL
    query := fmt.Sprintf(
        "INSERT INTO messages (%s) VALUES (%s)",
        strings.Join(cols, ", "),
        strings.Join(placeholders, ", "),
    )

    insertStmt, err := s.db.Prepare(query)
    if err != nil {
        return fmt.Errorf("erro ao preparar a declaração SQL: %w", err)
    }
    defer insertStmt.Close()

    // Passa a slice de interfaces para Exec
    _, err = insertStmt.Exec(values...)
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

