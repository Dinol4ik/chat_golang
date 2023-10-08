package main

import (
	"chat/internal/app"
	"chat/internal/app/config"
	"log"
	"os"
)

func main() {
	cfg, err := config.LoadConfig()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}
	application := app.New(cfg)
	err = application.Run()
	if err != nil {
		log.Println(err.Error())
		os.Exit(1)
	}

	os.Exit(0)
	//app := fiber.New()
	//
	//app.Use(cors.New())
	//var clients = make(map[*websocket.Conn]bool)
	//app.Use("/ws", func(c *fiber.Ctx) error {
	//	if websocket.IsWebSocketUpgrade(c) {
	//		c.Locals("allowed", true)
	//		return c.Next()
	//	}
	//	return fiber.ErrUpgradeRequired
	//})
	//
	//app.Get("/ws/:id", websocket.New(func(c *websocket.Conn) {
	//	clients[c] = true
	//	log.Println(c.Locals("allowed")) // true
	//	defer delete(clients, c)
	//
	//	var (
	//		mt  int
	//		msg []byte
	//		err error
	//	)
	//	for {
	//		if mt, msg, err = c.ReadMessage(); err != nil {
	//			log.Println("read:", err)
	//			break
	//		}
	//		log.Printf("recv: %s", msg)
	//		for conn := range clients {
	//			if err = conn.WriteMessage(mt, msg); err != nil {
	//				log.Println("write:", err)
	//				break
	//			}
	//		}
	//	}
	//
	//}))
	//
	//log.Fatal(app.Listen(":3001"))
}
