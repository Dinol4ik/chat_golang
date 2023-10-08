package http

import (
	"chat/internal/app/config"
	"chat/internal/usecase"
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"go.uber.org/zap"
)

type Server struct {
	config       *config.Config
	logger       *zap.SugaredLogger
	chatGolangUC usecase.ChatGolang
	connections  map[string]map[*websocket.Conn]string
}

func NewServer(config config.Config, logger zap.SugaredLogger, expUC usecase.ChatGolang, usersConn map[string]map[*websocket.Conn]string) *Server {
	return &Server{
		config:       &config,
		logger:       &logger,
		chatGolangUC: expUC,
		connections:  usersConn,
	}
}

func (s *Server) Run() error {
	app := fiber.New()
	app.Use(cors.New())
	app.Use("/ws", func(c *fiber.Ctx) error {
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})
	s.RegisterRoutes(app)
	return app.Listen(":8080")
}
func (s *Server) RegisterRoutes(app *fiber.App) {
	app.Get("/ws/:id", websocket.New(s.TryConnect))
	app.Get("/api/health-check", s.HealthCheck)
	app.Post("/api/createRooms", s.RoomCreating)
}
