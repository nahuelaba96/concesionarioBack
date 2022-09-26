package main

import (
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/ConcesionarioBack/router"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Panic("Error loading .env file")
	}

	router := router.NewRouter()
	puerto := os.Getenv("PORT")
	fmt.Println("Serve on port ", puerto)
	server := http.ListenAndServe(":"+puerto, router)
	log.Panic(server)
}
