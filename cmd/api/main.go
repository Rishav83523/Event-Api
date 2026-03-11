package main

import (
	"database/sql"
	"log"
	"rest-api/internal/database"
	"rest-api/internal/env"

	_ "github.com/joho/godotenv/autoload"
	_ "modernc.org/sqlite"
)

//autoload will automatically load the environment variables from the .env file when the application starts. This allows us to easily manage our configuration settings without hardcoding them into our code. We can access these environment variables using the env package we created earlier, which provides helper functions to retrieve string and integer values with default fallbacks.

type application struct { 
	port int
	jwtSecret string
	models  database.Models
}

func main() {
	db, err := sql.Open("sqlite", "./data.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	models := database.NewModels(db)

	app := &application{
		port : env.GetEnvInt("PORT", 8080),
		jwtSecret: env.GetEnvString("JWT_SECRET", "defaultsecret"),
		models: models,
	}

	if err := app.Serve(); err != nil {
		log.Fatal(err)
	}
}

