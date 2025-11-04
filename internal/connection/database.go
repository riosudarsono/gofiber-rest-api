package connection

import (
	"database/sql"
	"fmt"
	"gofiber-rest-api/internal/config"
	"log"

	_ "github.com/lib/pq"
)

func GetDatabase(conf config.Database) *sql.DB {
	dsn := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable timezone=%s",
		conf.Host,
		conf.Port,
		conf.User,
		conf.Pass,
		conf.Name,
		conf.Tz,
	)
	db, err := sql.Open("postgres", dsn)
	if err != nil {
		log.Fatal("Error connecting to open database", err.Error())
	}
	err = db.Ping()
	if err != nil {
		log.Fatal("Error connecting to ping database", err.Error())
	}
	return db
}
