package main

import (
	"encoding/json"
	"github.com/jackwhelpton/fasthttp-routing"
	"github.com/jackwhelpton/fasthttp-routing/access"
	"github.com/jackwhelpton/fasthttp-routing/content"
	"github.com/jackwhelpton/fasthttp-routing/fault"
	"github.com/jackwhelpton/fasthttp-routing/slash"
	"github.com/erikdubbelboer/fasthttp"
	"log"
)

type Customer struct {
	Firstname string `json:"firstname,omitempty"`
	Lastname  string `json:"lastname,omitempty"`
}

var customers []Customer

func main() {
	api := routing.New()
	customers = append(customers, Customer{Firstname: "Jean-Michel", Lastname: "Apeupres"})
	customers = append(customers, Customer{Firstname: "Michel", Lastname: "Platini"})

	api.Use(
		// all these handlers are shared by every route
		access.Logger(log.Printf),
		slash.Remover(fasthttp.StatusMovedPermanently),
		fault.Recovery(log.Printf),
	)

	api.Use(
		// these handlers are shared by the routes
		content.TypeNegotiator(content.JSON),
	)

	api.Get("/users", func(c *routing.Context) error {
		json.NewEncoder(c.Response.BodyWriter()).Encode(customers)
		return nil
	})

	fasthttp.ListenAndServe(":8080", api.HandleRequest)
}
