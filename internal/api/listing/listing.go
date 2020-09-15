package listing

import (
	"encoding/json"
	"github.com/mjrimrie/priceator/internal/datalayer"
	"github.com/mjrimrie/priceator/internal/resources/realtor"
	"net/http"
)

logger, _ := zap.NewProduction()

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
	db, err := datalayer.Connect()
	if err != nil{
		log)
	}
}

defer logger.Sync()
logger.Info("failed to fetch URL",
// Structured context as strongly typed Field values.
zap.String("url", url),
zap.Int("attempt", 3),
zap.Duration("backoff", time.Second),
)