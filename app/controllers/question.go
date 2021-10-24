package controllers

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
	"github.com/jungai/first-fiber-go/app/platform/database"
)

const (
	username = "host"
	password = "host"
	hostname = "127.0.0.1:3306"
	dbname   = "jungai_test"
)

type QuestionResult struct {
	Id       int    `json:"id"`
	Question string `json:"question"`
}

func getDsn() string {
	return fmt.Sprintf("%s:%s@tcp(%s)/%s", username, password, hostname, dbname)
}

func GetQuestions(c *fiber.Ctx) error {
	dsn := getDsn()
	db, err := database.DbCon(dsn)

	if err != nil {
		return fiber.ErrInternalServerError
	}

	var result = []*QuestionResult{}

	if result := db.Raw("Select * from Question").Scan(&result); result != nil {

		return c.JSON(result)
	}

	return fiber.ErrInternalServerError
}
