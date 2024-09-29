package server

import (
	"log"

	"github.com/gocql/gocql"
	"github.com/kataras/iris/v12"

	"go-user-crud-iris-cassandra/src/services/user"
)

// Server
type Server struct {
	address    string
	db_session *gocql.Session
}

// Return New Server
func NewServer(address string, db_session *gocql.Session) *Server {
	return &Server{address: address, db_session: db_session}
}

// Server RUN
func (s *Server) Run() error {
	// Fiber New App
	app := iris.New()

	// User Handler
	userStore := user.NewUserStore(s.db_session)
	userHandler := user.NewUserHandler(userStore)

	// Routes
	userHandler.RegisterRoutes(app)

	// Logs
	log.Println("Listening On", s.address)

	// Create Server
	return app.Listen(s.address)
}
