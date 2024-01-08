package search 

import (
	"errors"
	"fmt"
	"log"
	"net/url"
	"strconv"
	"strings"

	"github.com/schollz/closestmatch"
)

var (
	postalCodes = map[string]string{
		"Boat Quay" : "01", "Cecil" : "01", "Havelock Road": "01", "Marina": "01", "People’s Park": "01", "Raffles Place": "01", "Suntec City": "01", "Central": "01", "CBD": "01",
		"Anson": "02", "Chinatown": "02", "Shenton Way": "02", "Tanjong Pagar": "02",
		"Alexandra": "03", "Queenstown": "03", "Redhill": "03", "Tiong Bahru": "03",
		"Harbourfront": "04", "Keppel": "04", "Sentosa": "04", "Telok Blangah": "04",
		"Bouna Vista": "05", "Clementi": "05", "Dover": "05", "Hong Leong Garden": "05", "Pasir Panjang": "05", "West Coast": "05",
		"Beach Road (part)": "06", "City Hall": "06", "High Street": "06", "North Bridge Road": "06",
		"Beach Road": "07", "Bencoolen Road": "07", "Bugis": "07", "Golden Mile": "07", "Middle Road": "07", "Rocher": "07",
		"Farrer Park": "08", "Little India": "08", "Serangoon Road": "08",
		"Cairnhill": "09", "Killiney": "09", "Orchard": "09", "River Valley": "09",
		"Ardmore": "10", "Balmoral": "10", "Bukit Timah": "10", "Grange Road": "10", "Holland Road": "10", "Orchard Boulevard": "10", "Tanglin": "10",
		"Chancery ": "11", "Dunearn Road": "11", "Moulmein": "11", "Newton": "11", "Novena": "11", "Thomson": "11", "Watten Estate": "11",
		"Balestier": "12", "Toa Payoh": "12",
		"Braddell ": "13", "Macpherson": "13", "Potong Pasir": "13",
		"Eunos": "14", "Geylang": "14", "Kembangan": "14", "Paya Lebar": "14", "Sims": "14",
		"Amber Road": "15", "East Coast": "15", "Joo Chiat": "15", "Katong": "15", "Marine Parade": "15", "Meyer": "15", "Tanjong Rhu":" 15",
		"Bayshore": "16", "Bedok": "16", "Chai Chee": "16", "Eastwood": "16", "Kew Drive": "16", "Upper East Coast": "16",
		"Changi": "17", "Flora": "17", "Loyang": "17",
		"Pasir Ris": "18", "Simei": "18", "Tampines": "18",
		"Hougang": "19", "Ponggol": "19", "Sengkang": "19", "Serangoon Garden": "19", "Punggol" : "19",
		"Ang Mo Kio": "20", "Bishan": "20", "Braddell": "20",
		"Clementi Park": "21", "Hume Avenue": "21", "Ulu Pandan": "21", "Upper Bukit Timah":"21",
		"Boon Lay": "22", "Jurong": "22", "Tuas": "22",
		"Bukit Batok ": "23", "Bukit Panjang": "23", "Choa Chu Kang": "23", "Dairy Farm": "23", "Hillview": "23",
		"Lim Chu Kang": "24", "Sungei Gedong": "24", "Tengah": "24",
		"Admiralty Road": "25", "Kranji": "25", "Woodgrove": "25", "Woodlands": "25", 
		"Springleaf": "26", "Tagore": "26", "Upper Thomson": "26",
		"Admiralty Drive": "27", "Sembawang": "27", "Yishun":"27",
		"Seletar": "28", "Yio Chu Kang":"28",
	}
)

var (
	locationSet = []string{
		"Boat Quay", "Cecil", "Havelock Road", "Marina", "People’s Park", "Raffles Place", "Suntec City",
		"Anson", "Chinatown", "Shenton Way", "Tanjong Pagar",
		"Alexandra", "Queenstown", "Redhill", "Tiong Bahru",
		"Harbourfront", "Keppel", "Sentosa", "Telok Blangah",
		"Bouna Vista", "Clementi", "Dover", "Hong Leong Garden", "Pasir Panjang", "West Coast",
		"Beach Road (part)", "City Hall", "High Street", "North Bridge Road",
		"Beach Road", "Bencoolen Road", "Bugis", "Golden Mile", "Middle Road", "Rocher",
		"Farrer Park", "Little India", "Serangoon Road",
		"Cairnhill", "Killiney", "Orchard", "River Valley",
		"Ardmore", "Balmoral", "Bukit Timah", "Grange Road", "Holland Road", "Orchard Boulevard", "Tanglin",
		"Chancery", "Dunearn Road", "Moulmein", "Newton", "Novena", "Thomson", "Watten Estate",
		"Balestier", "Toa Payoh",
		"Braddell", "Macpherson", "Potong Pasir",
		"Eunos", "Geylang", "Kembangan", "Paya Lebar", "Sims",
		"Amber Road", "East Coast", "Joo Chiat", "Katong", "Marine Parade", "Meyer", "Tanjong Rhu",
		"Bayshore", "Bedok", "Chai Chee", "Eastwood", "Kew Drive", "Upper East Coast",
		"Changi", "Flora", "Loyang",
		"Pasir Ris", "Simei", "Tampines",
		"Hougang", "Ponggol", "Sengkang", "Serangoon Garden",
		"Ang Mo Kio", "Bishan", "Braddell",
		"Clementi Park", "Hume Avenue", "Ulu Pandan", "Upper Bukit Timah",
		"Boon Lay", "Jurong", "Tuas",
		"Bukit Batok", "Bukit Panjang", "Choa Chu Kang", "Dairy Farm", "Hillview",
		"Lim Chu Kang", "Sungei Gedong", "Tengah",
		"Admiralty Road", "Kranji", "Woodgrove", "Woodlands",
		"Springleaf", "Tagore", "Upper Thomson",
		"Admiralty Drive", "Sembawang", "Yishun",
		"Seletar", "Yio Chu Kang",
	}
)


type PriceType struct {
    	low int
	high int
}

func parsePriceRange(priceRange string) (PriceType, error) {

	values := strings.Split(priceRange, "-")
	min, err := strconv.Atoi(values[0])
        
        if err != nil {
                return PriceType{low: 0, high: 0}, fmt.Errorf("failed to parse min price: %v", err)
        }
        max, err := strconv.Atoi(values[1])
        if err != nil {
                return PriceType{low: 0, high:0}, fmt.Errorf("failed to parse max price: %v", err)
	}

	return PriceType{low: min, high: max}, nil
}

func parseInt(str string) (int, error) {
	num, err := strconv.Atoi(str)
	if err != nil {
		return 0, fmt.Errorf("failed to parse %s to integer: %v", str, err)
	}
	return num, nil
}

type MultiError struct {
	Errors []error
}

func (m MultiError) Error() string {
	var errorMessages []string
	for _, err := range m.Errors {
		if err != nil {
			errorMessages = append(errorMessages, err.Error())
		}
	}
	return strings.Join(errorMessages, "; ")
}

func mergeErrors(errors ...error) error {
	return MultiError{Errors: errors}
}

func parseUrlParam(queryParams url.Values) (string, PriceType, string, int, int, error) {

	location := queryParams.Get("location")
	priceRangeStr := queryParams.Get("priceRange")
	houseType := queryParams.Get("houseType")
	sizeStr := queryParams.Get("size")
	numberOfRoomsStr := queryParams.Get("numberOfRooms")

	priceRange, err_priceRange := parsePriceRange(priceRangeStr)
        size, err_size := parseInt(sizeStr)

        numberOfRooms, err_numberOfRooms := parseInt(numberOfRoomsStr)
 
        if priceRangeStr != "" && err_priceRange != nil {
                log.Println("[URL price range parsing error] : ", err_priceRange)
                priceRange = PriceType{low:0, high:100000}
        }
        if sizeStr != "" && err_size != nil {
                log.Println("[URL size parsing error] : ", err_size)
                return "", PriceType{low:0, high:0}, "Executive Condominium", 0, 0, errors.New("invalid size")
	}
        if numberOfRoomsStr != "" && err_numberOfRooms != nil {
                log.Println("[URL number of rooms parsing error] : ", err_numberOfRooms)
                return "", PriceType{low:0, high:0}, "Executive Condominium", 0, 0, errors.New("invalid number of rooms")
        }

	if priceRangeStr == "" {priceRange = PriceType{low: 0, high: 1000000}}
	if numberOfRoomsStr == "" {numberOfRooms = 9999}
	if numberOfRoomsStr == "NA" {numberOfRooms = 9999}
	if sizeStr == "" {size = -1}

	return location, priceRange, houseType, size, numberOfRooms, nil
}

func generateRangeString(size int) string {
	min := size
	max := size + 10

	rangeString := strconv.Itoa(min) + "-" + strconv.Itoa(max)
	return rangeString
}

func levenshteinMatchDistrictName(s string) (string) {
	log.Println("string to matched = ", s)
	bagSizes := []int{2}
	cm := closestmatch.New(locationSet, bagSizes)
	return cm.Closest(s)
}
