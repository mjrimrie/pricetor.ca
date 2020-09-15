package realtor

import (
	"encoding/json"
	"net/http"
	"net/url"
	"time"
)

type ErrorCode struct {
	Id          int
	Description string
	Status      string
	ProductName string
	Version     string
}
type SearchApiResponse struct {
	ErrorCode ErrorCode      `json:ErrorCode`
	Results   []SearchResult `json:Results`
}
type PropertyListingResponse struct {
	ErrorCode ErrorCode
	Property  Property
}
type SearchResult struct {
	Id                 int
	MlsNumber          int
	PublicRemarks      string
	Property           Property
	PostalCode         string
	RelativeDetailsURL string
	StatusId           int
	PhotoChangeDateUTC string
	Distance           int
	RelativeURLEn      string
	RelativeURLFr      string
}

type Property struct {
	Price   string
	Type    string
	Address Address
}
type Address struct {
	AddressText string
	Longitude float32
	Latitue float32
}


var searchDefaults = url.Values{"CultureId": {"1"}, "ApplicationId": {"1"}, "PropertySearchTypeId": {"1"}}
var listingDefaults = url.Values{"CultureId": {"1"}, "ApplicationId": {"37"}, "HashCode": {"0"}}
var httpClient = &http.Client{Timeout: 30 * time.Second}

const realtorSearchUrl = "https://api37.realtor.ca/Listing.svc/PropertySearch_Post"
const realtorListingUrl = "https://api37.realtor.ca//Listing.svc/PropertyDetails"

func getSearchDefaults() url.Values{
	defaults := url.Values{}
	for key, value := range searchDefaults {
		defaults.Add(key, value[0])
	}
	return defaults
}

func getListingDefaults() url.Values{
	defaults := url.Values{}
	for key, value := range listingDefaults {
		defaults.Add(key, value[0])
	}
	return defaults
}

func buildSearchParams(searchParams map[string]string) url.Values {
	var formParms = getSearchDefaults()
	if searchParams != nil && len(searchParams) > 0 {
		for key, value := range searchParams {
			formParms.Add(key, value)
		}
	}
	return formParms
}
func getListingUrlWithParams(listingId string, mlsNumber string, extra map[string]string) (*url.URL, error) {
	base, err := url.Parse(realtorListingUrl)
	if err != nil {
		return nil, err
	}
	params := buildListingParams(listingId, mlsNumber, extra)

	base.RawQuery = params.Encode()
	return base, nil
}
func buildListingParams(listingId string, mlsNumber string, extra map[string]string) url.Values{
	var queryParams = getListingDefaults()
	queryParams.Add("PropertyId", listingId)
	queryParams.Add("ReferenceNumber", mlsNumber)
	if extra != nil && len(extra) > 0{
		for key, value := range extra {
			queryParams.Add(key, value)
		}
	}
	return queryParams
}

func SearchListings(apiResponse interface{}, extra map[string]string) error{
	options := buildSearchParams(extra)
	resp, err := httpClient.PostForm(realtorSearchUrl, options)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&apiResponse)
	return nil
}

func GetPropertyListing(listingId string, mlsNumber string, details interface{}, extra map[string]string) error {
	url, err := getListingUrlWithParams(listingId, mlsNumber, nil)
	if err != nil{
		return err
	}
	resp, err := httpClient.Get(url.String())

	if err != nil{
		return err
	}
	defer resp.Body.Close()

	json.NewDecoder(resp.Body).Decode(&details)

	return nil
}