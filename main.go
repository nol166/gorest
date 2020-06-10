package main

import (
	"context"
	"encoding/json"

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
	collection, err := getMongoDbCollection(dbName, collectionName)

	if err != nil {
		c.Status(500).Send(err)
		return
	}

	var person Person
	json.Unmarshal([]byte(c.Body()), &person)

	res, err := collection.InsertOne(context.Background(), person)
	if err != nil {
		c.Status(500).Send(err)
		return
	}

	response, _ := json.Marshal(res)
	c.Send(response)
}

func updatePerson(c *fiber.Ctx) {

}

func deletePerson(c *fiber.Ctx) {

}
