package main

import (
	"flag"
	"strconv"

	"github.com/dasper/apiproxy"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"
)

type APIRequest struct {
	URL string `json:"url" form:"url" query:"url"`
}

func main() {
	app := fiber.New()

	app.Use(cors.New(cors.Config{
		AllowOrigins: "*",
		AllowMethods: "GET,POST",
		AllowHeaders: "Origin, Content-Type, Accept",
	}))

	host := flag.String("h", "0.0.0.0", "host")
	port := flag.Int("p", 8080, "port")

	addressString := *host + ":" + strconv.Itoa(*port)

	app.Get("/", func(c *fiber.Ctx) error {
		p := new(APIRequest)

		if err := c.QueryParser(p); err != nil {
			return err
		}

		response, err := apiproxy.GetResponse(p.URL)
		if err != nil {
			return err
		}

		c.Response().SetStatusCode(response.Code)
		c.Set("Content-type", response.Type)
		return c.SendString(response.Body)
	})

	err := app.Listen(addressString)
	if err != nil {
		return
	}
}
