package server

import (
	"fmt"
	"log"
	"net/http"
	"text/template"
)

type Server struct{}

func NewServer() *Server {
	return &Server{}
}

func (s *Server) Run() {
	http.HandleFunc("/", s.loginHandler)
	http.HandleFunc("/handle", s.Handle)

	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)

}

func (s *Server) loginHandler(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("./pkg/webui/login.html")
	if err != nil {
		log.Fatal(err)
	}

	t.Execute(w, nil)
}

func (s *Server) Handle(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("server handle func"))
}
