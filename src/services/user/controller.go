package user

import (
	"github.com/kataras/iris/v12"

	"go-user-crud-iris-cassandra/src/types"
)

func (h *userHandler) getUser(ctx iris.Context) {
}

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
		ctx.JSON(iris.Map{"error": "Failed to fetch users"})
		return
	}

	// JSON
	ctx.JSON(users)
}

func (h *userHandler) createUser(ctx iris.Context) {
}

func (h *userHandler) deleteUser(ctx iris.Context) {
}

func (h *userHandler) updateUser(ctx iris.Context) {
}
