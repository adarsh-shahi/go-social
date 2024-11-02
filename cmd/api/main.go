package main

import (
	"log"

	"github.com/adarsh-shahi/internal/env"
	"github.com/adarsh-shahi/internal/store"
)

func main() {
	cfg := config{

		//TODO: Get string env doesnt work currently it just uses the fallback
		addr: env.GetString("ADDR", ":8080"),
	}

	store := store.NewStorage(nil)
	app := &app{
		config: cfg,
		store: store,
	}

	mux := app.mount()
	log.Fatal(app.run(mux))
}
