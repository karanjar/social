package main

import (
	"log"
	"social/internal/db"
	"social/internal/env"
	"social/internal/store"
)

func main() {
	cfg := config{
		db: dbConfig{
			addr:         env.GetString("DB_ADDR", "postgres://admin:adminpassword@localhost/socialnetwork?sslmode=disable"),
			maxOPenconns: env.GetInt("DB_MAX_OPEN_CONNS", 30),
			maxIdleconns: env.GetInt("DB_MAX_IDLE_CONNS", 30),
			maxIdleTime:  env.GetString("DB_MAX_IDLE_TIME", "15m"),
		},
	}

	db, err := db.New(
		cfg.db.addr,
		cfg.db.maxOPenconns,
		cfg.db.maxIdleconns,
		cfg.db.maxIdleTime,
	)
	if err != nil {
		log.Panic(err)
	}
	defer db.Close()

	storage := store.NewPostgresStorage(db)
	log.Print("database connected successfully")

	app := &application{
		config: cfg,
		store:  storage,
	}
	mux := app.mount()

	log.Fatal(app.Run(mux))
}
