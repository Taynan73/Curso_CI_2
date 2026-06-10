package database

import (
	"log"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"fmt"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	
	stringDeConexao := fmt.Sprintf(
    "host=%s user=%s password=%s dbname=%s port=%s sslmode=disable",
    os.Getenv("HOST"),
    os.Getenv("USER"),
    os.Getenv("PASSWORD"),
    os.Getenv("DBNAME"),
    os.Getenv("DBPORT"),
)

	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}
