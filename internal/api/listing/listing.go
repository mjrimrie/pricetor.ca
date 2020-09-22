package listing

import (
	"encoding/json"
	"github.com/mjrimrie/priceator/internal/resources/realtor"
	"net/http"
)

type listingParams struct {
	listingId string
	mlsNumber string
}

func HandleListing(w http.ResponseWriter, req *http.Request){
	switch req.Method{
	case "GET":
		getListing(w, req)
		return
	default:
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}

}
func getListing(w http.ResponseWriter, req *http.Request) {
	listingId, listingIdExists := req.URL.Query()["listingId"]
	mlsNumber, mlsNumberExists := req.URL.Query()["mlsNumber"]
	err := false
	if !listingIdExists {
		http.Error(w, "listingId is None", http.StatusBadRequest)
		err = true
	}
	if !mlsNumberExists {
		http.Error(w, "mlsNumber is None", http.StatusBadRequest)
		err = true
	}
	if err == true {
		return
	}
	response := realtor.PropertyListingResponse{}
	realtor.GetPropertyListing(listingId[0], mlsNumber[0], &response, nil)

	if response.ErrorCode.Id > 200 {
		http.Error(w, response.ErrorCode.Description, response.ErrorCode.Id)
		return
	}
	json.NewEncoder(w).Encode(response.Property)

}
func AddListingToWatch(w http.ResponseWriter, req *http.Request) {
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}
}
