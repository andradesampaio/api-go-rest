package database

import ("gorm.io/gorm"
	"gorm.io/driver/postgres"
)

var (
	DB *gorm.DB
	err error
)

func Connect() {
	stringConexao := "host=localhost user=root password=root dbname=mydatabase port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(stringConexao))
	if err != nil {
		panic("failed to connect database")
	}
}