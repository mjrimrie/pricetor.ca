# Pricetor.ca

Add, track and receive pricing alarts on MLS properties listed on realtor.ca.


## API 

### /listing/watch/ -- POST
Add a listing for a user
POST request with content-type `application/json` with request body
```json
{
  "mlsNumber": "<MLS Number>",
  "listingId": "<Listing Id from realtor.ca URL>"
}
```

Returns 200 OK if successful with content-type `application/json` with body
```json
{
    "uid": "<generated uid>"
}
```

### /listings/ -- GET
Return 200 OK and all listings for user with content-type `application/json`
```json
[
   {
      "uid": "<generated uid>" str,
      "address": "<address of property>" str,
      "prices": [
        {
          "price": "<price>" int,
          "datetime": "<datetime of price>" UTC ISO-8601 
        },
        ... repeated per price
      ],     
      "mlsLink": "<link to realtor.ca page>" 
   },
   ...repeated per object
]


```
