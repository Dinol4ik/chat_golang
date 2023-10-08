package usecase

import (
	"chat/internal/params"
	"github.com/gofiber/fiber/v2"
)

type ChatGolang interface {
	ConnectClientsInRoom(user *params.UserParams) error
	CreateRoomsForLessons(ctx *fiber.Ctx, lessons params.Lessons, rooms params.Rooms) error
}
