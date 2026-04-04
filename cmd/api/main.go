package main

import (
	"fmt"
	"log"

	"github.com/rafael-dorneles/payment-gateway/internal/database"
)

func main() {
	fmt.Println("Iniciando ")

	// inicializa o banco
	dbPool, err := database.ConnectPostgres()
	if err != nil {
		log.Fatalf("Falha crítica ao conectar no banco: %v", err)
	}
	defer dbPool.Close()

	// rodas as migration
	if err := database.RunMigrations(); err != nil {
		log.Fatalf("Falha crítica ao conectar no banco: %v", err)
	}

	select {}
}
