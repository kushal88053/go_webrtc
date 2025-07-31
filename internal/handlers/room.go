package handlers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/websocket/v2"
	guuid "github.com/google/uuid"
)

func RoomCreate(c *fiber.Ctx) error {
	id := guuid.New().String()
	return c.Redirect(fmt.Sprintf("/room/%s", id), fiber.StatusFound)

}

func Room(c *fiber.Ctx) error {
	uuid := c.Params("uuid")
	if uuid == "" {
		return c.Status(fiber.StatusBadRequest).SendString("Invalid room UUID")
	}

	uuid, suuid, _ := createOrGetRoom(uuid)
}

func RoomWebsocket(c *websocket.Conn) {

	uuid := c.Params("uuid")

	if uuid == "" {
		c.Close(websocket.StatusBadRequest, "Invalid room UUID")
		return
	}

	_, _, room := createOrGetRoom(uuid)

}
