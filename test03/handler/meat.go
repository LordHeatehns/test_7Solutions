package handler

import (
	"strings"
	"test03/binding"
	"test03/service"

	"github.com/gofiber/fiber/v2"
)

func GetQuantityMeatHandler(c *fiber.Ctx) error {
	req := new(binding.Meat)
	if err := c.BodyParser(req); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(err)
	}
	text, err := service.GetText()
	if err != nil {
		panic(err)
	}

	splited := strings.Fields(text)
	data := make(map[string]int)

	for _, word := range splited {
		for _, item := range req.MeatList {
			if strings.EqualFold(word, item) {
				data[item] += 1
			}

		}
	}

	return c.Status(fiber.StatusOK).JSON(fiber.Map{
		"beef": data,
	})
}
