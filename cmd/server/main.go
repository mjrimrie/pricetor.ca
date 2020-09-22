package main

import (
	"github.com/mjrimrie/priceator/internal/datalayer"
	"github.com/mjrimrie/priceator/internal/server"
	"github.com/rs/zerolog"
	"github.com/rs/zerolog/log"
	"os"
)

func main() {
	database_url := os.Getenv("DATABASE_URL")
	if database_url == ""{
		log.Fatal().Msg("environment variable not set: DATABASE_URL")
		os.Exit(1)
	}
	log.Logger = log.Output(zerolog.ConsoleWriter{Out: os.Stderr})
	err := datalayer.Connect(database_url)
	if err != nil{
		log.Fatal().Err(err).Msg("unable to initialize datalayer")
		os.Exit(1)
	}
	log.Info().Msg("succesfully initialized datalayer")

	err = server.Start()
	if err != err{
		log.Fatal().Err(err).Msg("could not start server")
		os.Exit(1)
	}



}
