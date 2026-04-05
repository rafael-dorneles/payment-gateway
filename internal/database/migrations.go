package database

import (
	"errors"
	"fmt"
	"log"
	"os"

	"github.com/golang-migrate/migrate/v4"
	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
)

func RunMigrations() error {
	// 1. Tenta ler a URL completa do ambiente
	dbURL := os.Getenv("DB_URL")

	// 2. Se a variável estiver vazia (rodando local fora do docker), usa o padrão
	if dbURL == "" {
		dbURL = "postgres://user:password@localhost:5432/payment_db?sslmode=disable"
	}

	m, err := migrate.New(
		"file://migrations", // Verifique se esta pasta existe no seu container
		dbURL,
	)
	if err != nil {
		return fmt.Errorf("erro ao iniciar migration: %v", err)
	}

	if err := m.Up(); err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return fmt.Errorf("erro ao rodar migration up: %v", err)
	}

	log.Println("Migrations aplicadas com sucesso!")
	return nil
}
