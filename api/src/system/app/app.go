package app

import (
	"log"
	"net/http"

	"github.com/go-xorm/xorm"
	"github.com/gorilla/mux"
)

type Server struct {
	port string
	Db   *xorm.Engine
}

func NewServer() Server {
	return Server{}
}

func (s *Server) Init(port string, db *xorm.Engine) {
	log.Println("Initializing server....")
	s.port = ":" + port
	s.Db = db
}

func HomeHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))

}
func (s *Server) Start() {
	log.Println("Starting server on port " + s.port)

	r := mux.NewRouter()
	r.HandleFunc("/", HomeHandler)

	http.ListenAndServe(s.port, r)
}
