package main

import (
	"fmt"
	"io"
	"net/http"
	"os"

	"github.com/gofiber/fiber/v2"
)


const APP_PORT = ":4444"
const PATH = "public/upload"

func main(){

	router := fiber.New()

	router.Post("/", func(c *fiber.Ctx) error {
		username := c.FormValue("username")
		file, err := c.FormFile("file")
		if err != nil {
			return err
		}

		location := fmt.Sprintf("%s/%s", PATH, username)

		err = os.MkdirAll(location, os.ModePerm)
		if err != nil {
			return err
		}

		dst, err := os.Create(fmt.Sprintf("%s/%s",PATH, file.Filename ))
		if err != nil {
			return err
		}

		defer dst.Close()

		f, err := file.Open()
		
		if err != nil{
			return err
		}

		defer f.Close()

		if _, err := io.Copy(dst, f); err != nil {
			return err
		}

		return c.Status(http.StatusOK).JSON(fiber.Map{
			"messages" : "upload success",
		})


	})

	router.Listen(APP_PORT)
}