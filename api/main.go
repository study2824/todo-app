package main

import (
	"todo/api/db"
	"todo/api/router"
)

func main() {
	db.Init()
	router.Init()
}