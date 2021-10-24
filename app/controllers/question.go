package controllers

import (
	"github.com/gofiber/fiber/v2"
	"github.com/jungai/first-fiber-go/app/pkg/utils"
	"github.com/jungai/first-fiber-go/app/platform/database"
)

type QuestionResult struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
}

func GetQuestions(c *fiber.Ctx) error {
	db, err := database.DbCon(utils.GetDsn(&utils.JungaiTest))

	if err != nil {
		return fiber.ErrInternalServerError
	}

	var result = []QuestionResult{}

	if db.Raw("Select * from Question").Scan(&result); result != nil {

		return c.JSON(result)
	}

	return fiber.ErrInternalServerError
}
