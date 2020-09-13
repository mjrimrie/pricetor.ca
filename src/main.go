package main

import (
	"fmt"
	"resources/realtor"
	"server"
)

func main() {
	var listingResponse realtor.PropertyListingResponse
	var searchResponse realtor.SearchApiResponse

	realtor.SearchListings(&searchResponse, nil)
	fmt.Printf("Search Response Code: %d\n", searchResponse.ErrorCode.Id)

	realtor.GetPropertyListing("22354329", "855587", &listingResponse, nil)

	fmt.Printf("Listing Response Code: %d\n", listingResponse.ErrorCode.Id)

	server.Start()
}

