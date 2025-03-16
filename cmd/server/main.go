package main

import (
	"flag"
	"fmt"
	"net/http"
	_ "net/http/pprof"
	"os"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"github.com/sirupsen/logrus"
	"github.com/walterfan/go-mmuc/internal"
	"github.com/walterfan/go-mmuc/internal/config"
	"github.com/walterfan/go-mmuc/internal/message"
	"github.com/walterfan/go-mmuc/internal/service"
)

var (
	configPath string
	version    bool

	wsUpgrader = websocket.Upgrader{
		CheckOrigin: func(r *http.Request) bool { return true },
	}
)

func init() {
	flag.StringVar(&configPath, "f", "config/config.yaml", "path to the configuration file")
	flag.BoolVar(&version, "v", false, "print version")
	flag.BoolVar(&version, "version", false, "print version")
	flag.Parse()
}

func websocketHandler(c *gin.Context) {
	conn, err := wsUpgrader.Upgrade(c.Writer, c.Request, nil)
	if err != nil {
		logrus.Println("WebSocket upgrade failed:", err)
		return
	}
	defer conn.Close()

	for {
		var req message.Request
		if err := conn.ReadJSON(&req); err != nil {
			logrus.Println("Read error:", err)
			break
		}

		resp := message.Response{
			Request: req,
			Code:    200,
			Desc:    "Message received successfully",
		}

		if err := conn.WriteJSON(resp); err != nil {
			logrus.Println("Write error:", err)
			break
		}
	}
}

func main() {
	if version {
		internal.PrintVersion()
		return
	}

	// Load configuration
	cfg, err := config.LoadConfig(configPath)
	if err != nil {
		fmt.Printf("Failed to load config: %v", err)
		os.Exit(1)
	}
	fmt.Printf("Loaded config: %v\n", cfg)

	// Initialize logger
	err = config.InitLogger(cfg.LogDir, cfg.LogLevel)
	if err != nil {
		fmt.Printf("Failed to initialize logger: %v", err)
		os.Exit(1)
	}
	defer config.CloseLogger()

	r := gin.Default()

	user_service, err := service.NewUserService(cfg.DbUsername, cfg.DbPassword, cfg.DbHost, cfg.DbPort, cfg.DbName)
	if err != nil {
		logrus.Fatalf("Failed to initialize logger: %v", err)
	}

	r.POST("/users", user_service.CreateUser)
	r.GET("/users", user_service.ListUsers)
	r.GET("/users/:id", user_service.GetUser)
	r.PUT("/users/:id", user_service.UpdateUser)
	r.DELETE("/users/:id", user_service.DeleteUser)
	r.GET("/ws", websocketHandler)

	logrus.Println("Server started on :%d", cfg.ServerPort)
	r.Run(fmt.Sprintf(":%d", cfg.ServerPort))
}
