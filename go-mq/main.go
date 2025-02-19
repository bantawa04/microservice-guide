package main

import (
	"log"

	"github.com/bantawa04/go-mq/bootstrap"
)

// @title Fiber Boilerplate API
// @version 1.0
// @description This is a sample swagger for Fiber Boilerplate
// @termsOfService http://swagger.io/terms/
// @contact.name API Support
// @contact.email your-email@domain.com
// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html
// @host localhost:8080
// @BasePath /api
func main() {
	app := bootstrap.NewApplication()
	log.Fatal(app.Listen(":5001"))
}
