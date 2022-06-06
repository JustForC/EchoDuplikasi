package configs

import (
	"Kynesia/ent"
	"log"

	_ "github.com/go-sql-driver/mysql"
	"github.com/joho/godotenv"
)

func Connect() *ent.Client {
	env, errEnv := godotenv.Read()

	if errEnv != nil {
		log.Fatal("Error in reading env!")
	}

	dsn := env["DB_USERNAME"] + ":" + env["DB_PASSWORD"] + "@tcp(" + env["DB_HOST"] + ":" + env["DB_PORT"] + ")/" + env["DB_NAME"] + "?parseTime=True"
	client, errConnect := ent.Open(env["DB_DRIVER"], dsn)

	if errConnect != nil {
		log.Fatal(errConnect.Error())
	}

	return client
}
