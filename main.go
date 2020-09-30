package main

import (
		"log"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/template/html"
)

func main(){
	engine := html.New("./web/views",".html")
	engine.Reload(true)
	engine.Debug(true)
	app:= fiber.New(fiber.Config{
		Views:engine,
	})

	// app.Get("/main/:value",func(c *fiber.Ctx) error {
	// 	return c.Render("layout",fiber.Map{
	// 		"Title":"hai dunia macan",
	// 		"nama":"resiko",
	// 		"bolang":1234,
	// 	})
	// })
	app.Get("/main/:value", func(c *fiber.Ctx) error {
		// Render index within layouts/main
		return c.Render("layout", fiber.Map{
			"Title": "dunia",
			"nama":"ini nama",
			"bolang":"saya bolang",
		}, "layouts/main")
	})
log.Fatal(app.Listen(":3000"))

}