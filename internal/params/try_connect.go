package params

import (
	"github.com/gofiber/contrib/websocket"
)

type UserParams struct {
	UserConnection *websocket.Conn
	ListRooms      map[string]map[*websocket.Conn]string
	IdRoom         string
}
