package app

import (
	"chat/internal/app/config"
	"chat/internal/handlers/http"
	"chat/internal/storage/repo"
	chatGolang2 "chat/internal/usecase/chatGolang"
	"context"
	"fmt"
	"github.com/gofiber/contrib/websocket"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"os"
	"os/signal"
	"syscall"
)

type App struct {
	cfg *config.Config
}

func (a *App) Run() error {
	atom := zap.NewAtomicLevel()
	atom.SetLevel(zapcore.Level(*a.cfg.Logger.Level))
	encoderConfig := zap.NewProductionEncoderConfig()
	encoderConfig.EncodeTime = zapcore.RFC3339TimeEncoder
	zapCore := zapcore.NewCore(
		zapcore.NewJSONEncoder(encoderConfig),
		os.Stdout,
		atom,
	)
	logger := zap.New(zapCore)
	logger = logger.With(zap.String("service", "experiment"))
	log := logger.Sugar()
	atom.SetLevel(zapcore.Level(*a.cfg.Logger.Level))
	log.Infof("logger initialized successfully")
	ct := context.Background()
	opt := redis.NewClient(&redis.Options{Addr: a.cfg.Redis.Host, Password: a.cfg.Redis.Password, DB: a.cfg.Redis.DBName})
	repository := repo.NewConnect(opt)
	log.Info(opt.ClientInfo)
	fmt.Println(opt.Info(ct))
	chatGolang := chatGolang2.NewChatGolangUseCase(&repository, log)

	var conn = make(map[string]map[*websocket.Conn]string)
	httpServer := http.NewServer(*a.cfg, *log, chatGolang, conn)
	log.Info("application has started")
	go httpServer.Run()

	exit := make(chan os.Signal, 2)

	signal.Notify(exit, syscall.SIGINT, syscall.SIGTERM)

	<-exit

	log.Debug("waiting for httpServer to shut down")

	log.Info("application has been shut down")

	return nil
}

func New(cfg *config.Config) *App {
	return &App{cfg}
}
