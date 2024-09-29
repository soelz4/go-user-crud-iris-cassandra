package server

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/kataras/iris/v12"
)

// Server
type Server struct {
	address string
	db      *gocql.Session
}

// Return New Server
func NewServer(address string, db *gocql.Session) *Server {
	return &Server{address: address, db: db}
}

// Server RUN
func (s *Server) Run() error {
	// Fiber New App
	app := iris.New()

	// Handlers

	// Routes

	// Logs
	log.Println("Listening On", s.address)

	// Create Server
	return app.Listen(s.address)
}
