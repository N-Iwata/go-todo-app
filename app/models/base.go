package models

import (
	"crypto/sha1"
	"database/sql"
	"fmt"
	"log"

	"todo-app/config"

	"github.com/google/uuid"
	_ "github.com/lib/pq"
)

var Db *sql.DB
var err error

func init() {
	Db, err = sql.Open(config.Config.SQLDriver, fmt.Sprintf(`user=admin dbname=%s password=admin sslmode=disable`, config.Config.DbName))
	if err != nil {
		log.Panicln(err)
	}
	fmt.Println(Db)
}

func createUUID() (uuidobj uuid.UUID) {
	uuidobj, _ = uuid.NewUUID()
	return uuidobj
}

func Encrypt(plaintext string) (cryptext string) {
	cryptext = fmt.Sprintf("%x", sha1.Sum([]byte(plaintext)))
	return cryptext
}
