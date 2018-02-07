package main

import (
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {

	a := App{}

	a.InitializeApplication(os.Getenv("MYSQL_USERNAME"),
		os.Getenv("MYSQL_PASSWORD"),
		os.Getenv("MYSQL_DB"),
		os.Getenv("MYSQL_DB_URL"))

	a.RunApplication(":7070")
}
