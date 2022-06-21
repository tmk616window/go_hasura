package main

import (
	"api_test/models"
)


func main() {
	db := models.DbConnect()
	defer db.Close()
	db.LogMode(true)
}
