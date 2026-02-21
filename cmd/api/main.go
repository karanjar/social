package main

import "log"

func main() {
	cfg := config{
		addr: ":3015",
	}

	app := &application{
		config: cfg,
	}
	mux := app.mount()

	log.Fatal(app.Run(mux))
}
