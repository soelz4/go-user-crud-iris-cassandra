package user

import (
	"fmt"

	"github.com/gocql/gocql"
	"github.com/kataras/iris/v12"

	"go-user-crud-iris-cassandra/src/types"
)

// GetUser handles GET /users/{id}
func (h *userHandler) getUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userID, err := gocql.ParseUUID(id)
	// Error Handling
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid User ID"})
		return
	}

	var user types.User
	err = h.store.db_session.Query("SELECT * FROM users WHERE id = ? LIMIT 1", userID).
		Scan(&user.ID, &user.Name, &user.Email, &user.Phone)
	if err != nil {
		ctx.StatusCode(iris.StatusNotFound)
		ctx.JSON(iris.Map{"error": "User Not Found"})
		return
	}

	// JSON
	ctx.JSON(user)
}

// ListUsers handles GET /users
func (h *userHandler) listUsers(ctx iris.Context) {
	var users []types.User
	iter := h.store.db_session.Query("SELECT * FROM users").Iter()
	var user types.User

	for iter.Scan(&user.ID, &user.Name, &user.Email, &user.Phone) {
		users = append(users, user)
	}

	// Error Handling
	err := iter.Close()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": "Failed to Fetch Users"})
		return
	}

	// JSON
	ctx.JSON(users)
}

// CreateUser handles POST /users
func (h *userHandler) createUser(ctx iris.Context) {
	// Read JSON and Put into the User Structure
	var user types.User
	err := ctx.ReadJSON(&user)
	// Error Handling
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid Input"})
		return
	}

	user.ID = gocql.TimeUUID()

	err = h.store.db_session.Query("INSERT INTO users (id, name, email, phone) VALUES (?, ?, ?, ?)",
		user.ID, user.Name, user.Email, user.Phone).Exec()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.StatusCode(iris.StatusCreated)
	ctx.JSON(iris.Map{"message": fmt.Sprintf("User with ID=%v Created", user.ID)})
}

// DeleteUser handles DELETE /users/{id}
func (h *userHandler) deleteUser(ctx iris.Context) {
	id := ctx.Params().Get("id")
	userID, err := gocql.ParseUUID(id)
	// Error Handling
	if err != nil {
		ctx.StatusCode(iris.StatusBadRequest)
		ctx.JSON(iris.Map{"error": "Invalid User ID"})
		return
	}

	err = h.store.db_session.Query("DELETE FROM users WHERE id = ?", userID).Exec()
	if err != nil {
		ctx.StatusCode(iris.StatusInternalServerError)
		ctx.JSON(iris.Map{"error": err.Error()})
		return
	}

	ctx.JSON(iris.Map{"message": fmt.Sprintf("User with ID=%v Deleted", userID)})
}

func (h *userHandler) updateUser(ctx iris.Context) {
}
