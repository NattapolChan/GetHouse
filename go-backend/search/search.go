package search 

import (
	"log"
	"net/http"
	"encoding/json"
	"github.com/NattapolChan/ent"
	"github.com/NattapolChan/ent/propertylisting"
)

func IsValidHouseType(category string) bool {
    switch category {
    case
        "Detached House",
        "Non-landed Properties",
        "Semi-Detached House",
        "Terrace House",
	"Executive Condominium",
	"":
        return true
    }
    return false
}

func HandleSearchRequest(res http.ResponseWriter, req *http.Request, dbclient *ent.Client) {

	queryParams := req.URL.Query()
	
	location, priceRange, houseType, size, numberOfRooms, err := parseUrlParam(queryParams)

	if !IsValidHouseType(houseType) {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid HouseType"))
		return
	}

	if err != nil {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte(err.Error()))
		return
	} 

	numberOfRooms_lower := 0
	numberOfRooms_higher := 10000

	if numberOfRooms == 0 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid number of room (should not be zero)"))
		return
	} else if numberOfRooms <= -1 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid number of room (should not be negative)"))
		return
	} else if numberOfRooms == 9999 {

	} else {
		numberOfRooms_lower = numberOfRooms 
		numberOfRooms_higher = numberOfRooms 
	}

	matchedLocation := levenshteinMatchDistrictName(location)
	district := postalCodes[matchedLocation]

	size_lower := 0
	size_higher := 10000

	if size > 0 {
		size_lower = size - size/4
		size_higher = size + size/4
	}
	if size == 0 {
		res.WriteHeader(http.StatusBadRequest)
		res.Write([]byte("Invalid size"))
		return
	}

	if houseType == "" {
		u, _ := dbclient.PropertyListing.
			Query().
			Where(
				propertylisting.And(
					propertylisting.And(
						propertylisting.RentalPriceGTE(priceRange.low),
						propertylisting.RentalPriceLTE(priceRange.high),
					),
					propertylisting.And(
						propertylisting.NumOfBedroomGTE(numberOfRooms_lower),
						propertylisting.NumOfBedroomLTE(numberOfRooms_higher),
					),
					propertylisting.And(
						propertylisting.AreaGTE(size_lower),
						propertylisting.AreaLTE(size_higher),
					),
					propertylisting.Or(
						propertylisting.DistrictEQ(district),
						propertylisting.StreetContains(location),
						propertylisting.StreetHasPrefix(location),
					),
				),
			).Limit(200).All(req.Context())
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(u)

	} else {
		u, _ := dbclient.PropertyListing.Query().
			Where(
				propertylisting.And(
					propertylisting.And(
						propertylisting.RentalPriceGTE(priceRange.low),
						propertylisting.RentalPriceLTE(priceRange.high),
					),
					propertylisting.And(
						propertylisting.NumOfBedroomGTE(numberOfRooms_lower),
						propertylisting.NumOfBedroomLTE(numberOfRooms_higher),
					),
					propertylisting.And(
						propertylisting.AreaGTE(size_lower),
						propertylisting.AreaLTE(size_higher),
					),
					propertylisting.Or(
						propertylisting.DistrictEQ(district),
						propertylisting.StreetContains(location),
						propertylisting.StreetHasPrefix(location),
					),
					propertylisting.HouseTypeEQ(houseType),
				),
			).
			Limit(200).
			All(req.Context())
		res.Header().Set("Content-Type", "application/json")
		res.WriteHeader(http.StatusOK)
		json.NewEncoder(res).Encode(u)
	}
	
	log.Println(location)
}
