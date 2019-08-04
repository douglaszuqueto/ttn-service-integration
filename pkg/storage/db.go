package storage

import (
	"database/sql"
	"fmt"
	"log"
	"time"

	"github.com/douglaszuqueto/ttn-service-integration/config"
	"github.com/lib/pq"
	"github.com/pkg/errors"
)

var db *sql.DB

var (
	errItemNotFound = errors.New("Item não encontrado")
)

// GetConn getConn
func GetConn() *sql.DB {
	return db
}

// Connect init connection
func Connect() {
	cfg := config.Cfg

	connStr := fmt.Sprintf("postgres://%s:%s@%s/%s?sslmode=disable",
		cfg.Postgres.User,
		cfg.Postgres.Password,
		cfg.Postgres.Host,
		cfg.Postgres.Database,
	)

	log.Println("[BD] Iniciando conexão com banco de dados")

	var err error
	db, err = sql.Open("postgres", connStr)

	db.SetMaxIdleConns(5)
	db.SetMaxOpenConns(5)
	db.SetConnMaxLifetime(5 * time.Minute)

	if err != nil {
		log.Fatal(err)
	}

	ping()

	log.Println("[BD] Iniciado")
}

func ping() {
	log.Println("[DB] - Ping")
	err := db.Ping()
	checkErr(err)
	log.Println("[DB] - Pong")
}

// CloseConnection closeConnection
func CloseConnection() {
	log.Println("[BD] Fechando conexão...")
	db.Close()
	log.Println("[BD] Conexão fechada!")
}

func checkErr(err error) {
	if err != nil {
		log.Fatal(err)
	}
}

var (
	ErrAlreadyExists             = errors.New("object already exists")
	ErrDoesNotExist              = errors.New("object does not exist")
	ErrUsedByOtherObjects        = errors.New("this object is used by other objects, remove them first")
	ErrUserInvalidUsername       = errors.New("username name may only be composed of upper and lower case characters and digits")
	ErrUserPasswordLength        = errors.New("passwords must be at least 6 characters long")
	ErrInvalidUsernameOrPassword = errors.New("invalid username or password")
	ErrOrganizationInvalidName   = errors.New("invalid organization name")
	ErrInvalidEmail              = errors.New("invalid e-mail")
)

// HandlePSQLError handlePSQLError
// https://github.com/brocaar/lora-app-server/blob/master/internal/storage/errors.go#L41
func HandlePSQLError(err error) error {
	if err == sql.ErrNoRows {
		return ErrDoesNotExist
	}

	switch err := err.(type) {
	case *pq.Error:
		switch err.Code.Name() {
		case "unique_violation":
			return errors.Wrap(ErrAlreadyExists, err.Constraint)
		case "foreign_key_violation":
			return ErrDoesNotExist
		default:
			log.Println(err.Code.Name())
			return err
		}
	}

	return err
}

func doInsert(query string, args ...interface{}) (string, error) {
	stmt, err := db.Prepare(query)
	if err != nil {
		return "", err
	}
	defer stmt.Close()

	var id string

	err = stmt.QueryRow(args...).Scan(&id)
	if err != nil {
		return "", HandlePSQLError(err)
	}

	return id, nil
}
