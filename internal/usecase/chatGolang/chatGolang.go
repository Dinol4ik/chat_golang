package chatGolang

import (
	"chat/internal/params"
	"chat/internal/storage"
	"context"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"go.uber.org/zap"
	"log"
)

type UseCase struct {
	repo   storage.Repo
	logger *zap.SugaredLogger
}

func NewChatGolangUseCase(r storage.Repo, logger *zap.SugaredLogger) *UseCase {
	return &UseCase{
		repo:   r,
		logger: logger,
	}
}

func (uc *UseCase) ConnectClientsInRoom(user *params.UserParams) error {
	c := user.UserConnection
	s := user.ListRooms
	if s[user.IdRoom] != nil {
		roomNumber := s[user.IdRoom]
		roomNumber[c] = c.Query("connectId")
	} else {
		log.Println("Такой комнаты не существует!")
	}
	defer delete(s[user.IdRoom], c)
	ctx := context.Background()
	var (
		mt       int
		msg      []byte
		err      error
		messages []string
	)
	log.Println("Ошибка не в базе")
	messages, err = uc.repo.TakeLastMessages(ctx, user.IdRoom)
	if err != nil {
		log.Println("ошибка в базе")
	}
	if err = c.WriteJSON(messages); err != nil {
		log.Println("write:", err)
	}
	log.Println(s)
	for {
		if mt, msg, err = c.ReadMessage(); err != nil {

			log.Println("read:", err)
			break
		}
		err = uc.repo.AddMessage(ctx, string(msg), user.IdRoom)
		if err != nil {
			return err
		}
		log.Printf("recv: %s", msg)
		_ = mt
		for conn := range s[user.IdRoom] {
			if err = conn.WriteJSON(string(msg)); err != nil {
				log.Println("write:", err)
				break
			}

		}
	}
	return nil
}
func (uc *UseCase) CreateRoomsForLessons(ctx *fiber.Ctx, lessons params.Lessons, rooms params.Rooms) error {
	serverRooms := rooms.ChatRooms
	for _, i := range lessons.LessonsId {
		serverRooms[i] = make(map[*websocket.Conn]string)
		//serverRooms[i] = make(map[*websocket.Conn]string)
	}
	log.Println(rooms)
	return nil
}
