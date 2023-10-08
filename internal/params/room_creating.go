package params

import "github.com/gofiber/contrib/websocket"

type Rooms struct {
	ChatRooms map[string]map[*websocket.Conn]string
}
type Lessons struct {
	LessonsId []string `json:"lessonsId"`
}
