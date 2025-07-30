package server

import (
	"flag"
	"os"
	"time"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
	"github.com/gofiber/fiber/v2/middleware/logger"
	"github.com/gofiber/template/html"

	"your_project/handlers"
	"your_project/webSocket"
	"your_project/webrtc"
)

var (
	address = flag.String("addr", ":"+os.Getenv("PORT"), "server address")
	cert    = flag.String("cert", "", "TLS certificate file")
	key     = flag.String("key", "", "TLS key file")
)

func Run() error {
	flag.Parse()

	if *address == ":" || *address == ":"+"" {
		*address = ":8080"
	}

	engine := html.New("./views", ".html")
	app := fiber.New(fiber.Config{
		Views:                 engine,
		DisableStartupMessage: true,
	})

	app.Use(logger.New())
	app.Use(cors.New(cors.Config{
		AllowOrigins:     "*",
		AllowMethods:     "GET,POST,PUT,DELETE,OPTIONS",
		AllowHeaders:     "Origin, Content-Type, Accept, Authorization",
		ExposeHeaders:    "Content-Length, Access-Control-Allow-Origin, Access-Control-Allow-Headers",
		AllowCredentials: true,
	}))

	// Routes
	app.Get("/", handlers.Welcome)

	app.Get("/room/:id", handlers.RoomCreate)
	app.Get("/room/:uuid", handlers.Room)

	app.Get("/room/:uuid/websocket", webSocket.New(handlers.RoomWebsocket, webSocket.Config{
		HandShakeTimeout: 10 * time.Second,
	}))

	app.Get("/room/:uuid/chat", handlers.RoomChat)
	app.Get("/room/:uuid/chat/websocket", webrtc.New(handlers.RoomChatWebsocket))
	app.Get("/room/:uuid/viewer/websocket", webrtc.New(handlers.RoomViewerWebsocket))

	app.Get("/stream/:ssuid", handlers.Stream)
	app.Get("/stream/:ssuid/websocket", handlers.StreamWebsocket)
	app.Get("/stream/:ssuid/chat/websocket", handlers.StreamChatWebsocket)
	app.Get("/stream/:ssuid/viewer/websocket", webrtc.New(handlers.StreamViewerWebsocket))

	return app.Listen(*address)
}
