package main

import (
	"fmt"
	"github.com/gofiber/fiber/v2"
	"github.com/google/uuid"
	"ioc/di"
)

type PizzaService struct {
	Id uuid.UUID
}

func main() {
	i := 1
	app := fiber.New()

	di.Add("pizzaSrv", di.Singleton, func() *PizzaService {
		return &PizzaService{Id: uuid.New()}
	})

	/* Or could be done like this. Although empty structs are 0 bytes in size, I don't think this is a good idea lol.

	di.Add(PizzaService{}, di.Singleton, func()*PizzaService {
		return &PizzaService{Id: uuid.New()}
	})

	*/

	app.Get("/", func(c *fiber.Ctx) error {
		pizzaSrv := di.Resolve[*PizzaService]("pizzaSrv")

		// Again, this could be done like this if registered with the respective struct as the key:
		//pizzaSrv := di.Resolve[*PizzaService](PizzaService{})

		fmt.Println("Request:", i, pizzaSrv.Id)
		i++

		return nil
	})

	if err := app.Listen(":3000"); err != nil {
		panic(err)
	}

}
