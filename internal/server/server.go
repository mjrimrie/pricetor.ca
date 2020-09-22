package server

import (
	"github.com/mjrimrie/priceator/internal/api/listing"
	"github.com/rs/zerolog/log"
	"net/http"
)



func Start() error {
	http.HandleFunc("/listing/", listing.HandleListing)
	http.HandleFunc("/listing/watch/", listing.AddListingToWatch)

	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Error().Err(err).Msg("unable to start HTTP server")
		return err
	}
	log.Info().Msg("succesfully started HTTP server")
	return nil
}
