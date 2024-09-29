package user

import (
	"github.com/kataras/iris/v12"
)

// User Handler
type userHandler struct {
	store *userStore
}

// Return New User Handler
func NewUserHandler(store *userStore) *userHandler {
	return &userHandler{store: store}
}

// Routes (API Endpoints)
func (h *userHandler) RegisterRoutes(app *iris.Application) {
	app.Get("/users/{id:uuid}", h.getUser)
	app.Get("/users", h.listUsers)
	app.Post("/users", h.createUser)
	app.Delete("/users/{id:uuid}", h.deleteUser)
	app.Put("/users/{id:uuid}", h.updateUser)
}
