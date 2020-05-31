package config

import (
	"database/sql"
	_ "github.com/mattn/go-sqlite3"
	"os"
)

type Configuration struct {
	DatabaseInfo *sql.DB
	Err          error
}

const driverName = "sqlite3"

const recipeTable = "CREATE TABLE IF NOT EXISTS recipe_information(id INTEGER PRIMARY KEY,recipe_name TEXT ,desc TEXT ,ingredients BLOB,preparation_steps BLOB,categories BLOB,tags BLOB)"

func New(dataSourceName string) *Configuration {
	os.Remove(dataSourceName)
	database, err := sql.Open(driverName, dataSourceName)
	statement, _ := database.Prepare(recipeTable)
	statement.Exec()
	var configuration = &Configuration{}
	configuration.DatabaseInfo = database
	configuration.Err = err
	return configuration
}
