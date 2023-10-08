package http

import (
	"chat/internal/params"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
)

func (s *Server) TryConnect(c *websocket.Conn) {
	userParams := &params.UserParams{
		UserConnection: c,
		ListRooms:      s.connections,
		IdRoom:         c.Params("id"),
	}
	err := s.chatGolangUC.ConnectClientsInRoom(userParams)
	if err != nil {
		s.logger.Error("%s", "Ошибка в подключении")
	}
}
func (s *Server) HealthCheck(ctx *fiber.Ctx) error {
	return ctx.SendStatus(fiber.StatusOK)
}
func (s *Server) RoomCreating(ctx *fiber.Ctx) error {
	chatParams := params.Rooms{ChatRooms: s.connections}
	paramsIdRoom := params.Lessons{}
	err := ctx.BodyParser(&paramsIdRoom)
	if err != nil {
		return err
	}
	err = s.chatGolangUC.CreateRoomsForLessons(ctx, paramsIdRoom, chatParams)
	if err != nil {
		return err
	}
	return ctx.SendStatus(fiber.StatusOK)
}
