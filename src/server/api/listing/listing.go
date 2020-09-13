package listing

import (
	"encoding/json"
	"net/http"
	"resources/realtor"
)

type listingParams struct {
	listingId string
	mlsNumber string
}

func GetListing(w http.ResponseWriter, req *http.Request) {
	if req.Method != "GET" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
		return
	}
	listingId, listingIdExists := req.URL.Query()["listingId"]
	mlsNumber, mlsNumberExists := req.URL.Query()["mlsNumber"]
	error := false
	if !listingIdExists {
		http.Error(w, "listingId is None", http.StatusBadRequest)
		error = true
	}
	if !mlsNumberExists {
		http.Error(w, "mlsNumber is None", http.StatusBadRequest)
		error = true
	}
	if error == true {
		return
	}
	response := realtor.PropertyListingResponse{}
	realtor.GetPropertyListing(listingId[0], mlsNumber[0], &response, nil)

	if response.ErrorCode.Id > 200{
		http.Error(w, response.ErrorCode.Description, response.ErrorCode.Id)
		return
	}
	json.NewEncoder(w).Encode(response.Property)


}
func AddListingToWatch(w http.ResponseWriter, req *http.Request){
	if req.Method != "POST" {
		http.Error(w, http.StatusText(http.StatusMethodNotAllowed), http.StatusMethodNotAllowed)
	}



}
