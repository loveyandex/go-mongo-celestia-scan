package util

import (
	"github.com/gofiber/fiber/v2"
	"github.com/loveyandex/go-mongo-celestia-scan/util/payload"
)

func Jackson(c *fiber.Ctx, j interface{}, e error) error {
	if e != nil {
		return e
	}
	return c.JSON(payload.ApiResponse{Status: true, Data: j})
}

func JacksonList(c *fiber.Ctx, j interface{}, tot int64, e error) error {
	if e != nil {
		return e
	}
	return c.JSON(payload.ApiListResponse{Status: true, Data: j, TotalCount: tot})
}
