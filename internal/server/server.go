package server

import (
	"github.com/mjrimrie/priceator/internal/api/listing"
	"net/http"
)



func Start() {
	http.HandleFunc("/listing/", listing.GetListing)
	http.HandleFunc("/listing/watch/", listing.AddListingToWatch)
	http.ListenAndServe(":8080", nil)
}
