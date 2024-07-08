package bd

import (
	"log"
	"os"

	"github.com/leMoreira/api_simples/cmd/api/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaBaseDeDados() {

	stringDeConexao := os.Getenv("DATABASE")

	DB, err = gorm.Open(postgres.Open(stringDeConexao))

	if err != nil {
		log.Panic("Erro ao conectar ao banco de dados")
	}

	DB.AutoMigrate(
		&models.Task{},
		&models.User{},
		&models.AuthInput{})
}
