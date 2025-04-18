package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/javiermedinaj/uala-challenge/internal/handlers"
	"github.com/javiermedinaj/uala-challenge/internal/repository/sqlite"
	_ "github.com/mattn/go-sqlite3"
)

func main() {
	//iniciar la db
	db, err := sql.Open("sqlite3", "./uala.db")
	if err != nil {
		log.Fatal("error al abrir la db ", err)
	}
	defer db.Close()

	// Crea si no existen
	_, err = db.Exec(`
    CREATE TABLE IF NOT EXISTS users (
        id TEXT PRIMARY KEY,
        username TEXT NOT NULL UNIQUE,
        created_at DATETIME,
        updated_at DATETIME
    );
    CREATE TABLE IF NOT EXISTS tweets (
        id TEXT PRIMARY KEY,
        user_id TEXT,
        content TEXT,
        created_at DATETIME,
        updated_at DATETIME,
        FOREIGN KEY(user_id) REFERENCES users(id)
    );`)

	if err != nil {
		log.Fatal("error al crear las tablas ", err)
	}
	//inicializar el repositorio
	userRepo := sqlite.NewUserRepository(db)
	tweetRepo := sqlite.NewTweetRepository(db)

	//inicializar el handler
	h := handlers.NewHandler(userRepo, tweetRepo)

	//configuracion de gin
	r := gin.Default()

	//routes ejemplos post
	r.POST("/users", h.CreateUser)
	r.POST("/tweets", h.CreateTweet)

	//routes ejemplos get
	r.GET("/users/:id", h.GetUserByID)
	r.GET("/tweets/:id", h.GetTweetByID)

	//iniciar el server
	if err := r.Run(":8080"); err != nil && err != http.ErrServerClosed {
		log.Fatal("error al iniciar el server ", err)
	}

}
