package main

import (
	"flag"
	"go-vuejs-web-app/api/src/system/app"
	DB "go-vuejs-web-app/api/src/system/db"
	"os"

	"github.com/joho/godotenv"
)

var port string
var dbhost string
var dbport string
var dbuser string
var dbpass string
var dboptions string
var dbdatabase string

func init() {
	flag.StringVar(&port, "port", "8000", "Asing the port for the server")
	flag.Parse()

	if err := godotenv.Load("config.ini"); err != nil {
		panic(err)
	}

	envPort := os.Getenv("PORT")

	if len(envPort) > 0 {
		port = envPort
	}
}

func main() {
	db, err := DB.Connect()

	if err != nil {
		panic(err)
	}
	s := app.NewServer()
	s.Init(port, db)
	s.Start()
}
