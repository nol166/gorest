package main

import (
	"github.com/gofiber/fiber"
)

func main() {
	app := fiber.New()

	app.Get("/person", getPersons)
	app.Get("/person/:id", getPerson)
	app.Post("/person", createPerson)
	app.Put("/person/:id", updatePerson)
	app.Delete("/person/:id", deletePerson)

	app.Listen(8000)
}

const dbName = "personsDb"
const collectionName = "person"
const port = 8000

func getPersons(c *fiber.Ctx) {

}

func getPerson(c *fiber.Ctx) {

}

func createPerson(c *fiber.Ctx) {

}

func updatePerson(c *fiber.Ctx) {

}

func deletePerson(c *fiber.Ctx) {

}
