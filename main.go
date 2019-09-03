package main

import (
	"flag"
	"go-vuejs-web-app/src/system/app"
	"os"

	"github.com/joho/godotenv"
)

var port string

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
	s := app.NewServer()
	s.Init(port)
	s.Start()
}
