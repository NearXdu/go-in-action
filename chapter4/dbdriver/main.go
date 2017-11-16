package main

import (
	"database/sql"
	_ "zxcode/go-in-action/chapter4/dbdriver/postgres"
)

func main() {
	sql.Open("postgres", "mydb")
}
