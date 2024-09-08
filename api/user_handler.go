package api

import (
	"context"

	"github.com/gofiber/fiber/v3"
	"github.com/kaiquecaires/hotel_reservation/db"
	"github.com/kaiquecaires/hotel_reservation/types"
)

type UserHandler struct {
	userStore db.UserStore
}

func NewUserHandler(userStore db.UserStore) *UserHandler {
	return &UserHandler{
		userStore: userStore,
	}
}

func (h *UserHandler) HandleGetUser(c fiber.Ctx) error {
	var (
		id  = c.Params("id")
		ctx = context.Background()
	)
	user, err := h.userStore.GetUserById(ctx, id)
	if err != nil {
		return err
	}
	return c.JSON(user)
}

func (h *UserHandler) HandleGetUsers(c fiber.Ctx) error {
	u := types.User{
		FirstName: "James",
		LastName:  "Water cooler",
	}
	return c.JSON(u)
}
