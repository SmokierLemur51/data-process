package db

import (
	"fmt"
	"log"
	"os"

	"github.com/jmoiron/sqlx"
)

func CreateSchema(db *sqlx.DB, schemaPath string) error {
	// Using working directory
	data, err := os.ReadFile(schemaPath)
	if err != nil {
		return err
	}
	db.MustExec(string(data))
	log.Printf("\nDatabase <%s%s> populated\n", os.Getenv("DB_PATH"), os.Getenv("SQLITE_DB"))
	return nil
}

// No parameters, we use the environment variables
func CreateDatabaseFile() error {
	file, err := os.Create(
		fmt.Sprintf("%s/%s.db", os.Getenv("DB_PATH"), os.Getenv("SQLITE_DB")))
	if err != nil {
		return err
	}
	file.Close()
	log.Printf("\nDatabase <%s%s> created\n", os.Getenv("DB_PATH"), os.Getenv("SQLITE_DB"))
	return nil
}

// Delete then recreate database file,
func DeleteDatabase(db *sqlx.DB) error {
	// Delete file
	err := os.Remove(
		fmt.Sprintf("%s%s.db", os.Getenv("DB_PATH"), os.Getenv("SQLITE_DB")))
	if err != nil {
		return err
	}
	log.Printf("\nDatabase <%s%s> deleted\n", os.Getenv("DB_PATH"), os.Getenv("SQLITE_DB"))
	// Create new database file with existing function
	if err := CreateDatabaseFile(); err != nil {
		return err
	}
	return nil
}
