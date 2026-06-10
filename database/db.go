package database

import (
	"log"
	"strings"
	"os"
	"github.com/guilhermeonrails/api-go-gin/models"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var (
	DB  *gorm.DB
	err error
)

func ConectaComBancoDeDados() {
	
	// O strings.TrimSpace remove qualquer \n, \r ou espaço invisível do começo e do fim
	host     := strings.TrimSpace(os.Getenv("HOST"))
	user     := strings.TrimSpace(os.Getenv("USER"))
	password := strings.TrimSpace(os.Getenv("PASSWORD"))
	dbname   := strings.TrimSpace(os.Getenv("DBNAME"))
	port     := strings.TrimSpace(os.Getenv("DBPORT"))

	stringDeConexao := "host=" + host + " user=" + user + " password=" + password + " dbname=" + dbname + " port=" + port + " sslmode=disable"
	
	DB, err = gorm.Open(postgres.Open(stringDeConexao))
	if err != nil {
		log.Panic("Erro ao conectar com banco de dados")
	}

	DB.AutoMigrate(&models.Aluno{})
}