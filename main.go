package main

import (
	"database/sql"
	"os"

	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"

	_ "github.com/lib/pq"
	"github.com/wizlif/dfcu_bank/api"
	db "github.com/wizlif/dfcu_bank/db/sqlc"
	_ "github.com/wizlif/dfcu_bank/docs"
	_ "github.com/wizlif/dfcu_bank/docs/statik"
	"github.com/wizlif/dfcu_bank/util"
)

// @title DFCU Loan Service API
// @description The API can be used to maintain and log loan actions
func main() {
	config, err := util.LoadConfig(".")
	if err != nil {
		log.Fatal().Msg("Cannot load config")
	}

	if config.Environment == "development" {
		log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	}

	conn, err := sql.Open(config.DbDriver, config.DbSource)

	if err != nil {
		log.Fatal().Msg("Cannot connect to db")
	}

	store := db.NewStore(conn)
	runGinServer(config, store)
}

func runGinServer(config util.Config, store db.Store) {
	server, err := api.NewServer(config, store)

	if err != nil {
		log.Fatal().Err(err).Msg("could not start server")
	}

	err = server.Start(config.HTTPServerAddress)

	if err != nil {
		log.Fatal().Err(err).Str("address", config.HTTPServerAddress).Msg("Cannot start server ")
	}
}
