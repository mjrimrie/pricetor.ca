package server

import (
	"net/http"
	"server/api/listing"
)

func Start(){
	http.HandleFunc("/listing/", listing.GetListing)
	http.HandleFunc("/listing/watch/", listing.AddListingToWatch)
	http.ListenAndServe(":8080", nil)
}
