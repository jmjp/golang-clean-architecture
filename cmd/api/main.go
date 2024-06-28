package main

import (
	"onion/config"
	"onion/internal/infrastructure/database"
	"onion/internal/presentation/web"
)

func init() {
	_, err := config.Load()
	if err != nil {
		panic(err)
	}
}

func main() {
	db := database.NewPostgresDB()
	defer db.Close()
	web.NewServer(8080).Start(db)
}
