package schedule

import (
	"encoding/json"
	"github.com/NattapolChan/ent"
	"os"
	"net/http"
	"log"
	"fmt"
	"io"
	"context"
	"strconv"
	"strings"
	"github.com/NattapolChan/ent/propertylisting"
)

type Rental struct {
	AreaSqm      string `json:"areaSqm"`
	LeaseDate    string `json:"leaseDate"`
	PropertyType string `json:"propertyType"`
	District     string `json:"district"`
	AreaSqft     string `json:"areaSqft"`
	NoOfBedRoom  string `json:"noOfBedRoom"`
	Rent         int    `json:"rent"`
}

type Entry struct {
	Street  string   `json:"street"`
	X       string   `json:"x"`
	Project string   `json:"project"`
	Y       string   `json:"y"`
	Rental  []Rental `json:"rental"`
}

func UpdatePropertyListing(dbclient *ent.Client, url string) {
	fmt.Println("Update propertylisting")
	client := &http.Client{}
	req_ura, err_ura := http.NewRequest("GET", url, nil)
	req_ura.Header.Set("AccessKey", os.Getenv("URA_ACCESS_KEY"))
	req_ura.Header.Set("Token", os.Getenv("URA_API_TOKEN"))
	res_ura, _ := client.Do(req_ura)
	if err_ura != nil {
		log.Fatalln(err_ura)
	}
	fmt.Println(res_ura)
	body, err := io.ReadAll(res_ura.Body)
	if err != nil {
		log.Fatalln(err)
	}
	res_ura.Body.Close()
	sb := string(body)
	ReadPropertyListing(sb, dbclient)
}

func ReadPropertyListing(s string, dbclient *ent.Client) {

	fmt.Println(s)

	var outer struct {
		Status string `json:"Status"`
		Result []struct {
			Street  string `json:"street"`
			Project string `json:"project"`
			X       string `json:"x"`
			Y       string `json:"y"`
			Rental  []struct {
				AreaSqm    string `json:"areaSqm"`
				LeaseDate  string `json:"leaseDate"`
				PropertyType string `json:"propertyType"`
				District   string `json:"district"`
				AreaSqft   string `json:"areaSqft"`
				NoOfBedRoom string `json:"noOfBedRoom"`
				Rent       int    `json:"rent"`
			} `json:"rental"`
		} `json:"Result"`
	}

	err := json.Unmarshal([]byte(s), &outer)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}

	for _, item := range outer.Result {

		x, err := strconv.ParseFloat(item.X, 32)
		if err != nil {
			fmt.Println("Error : ", err)
			continue
		}
		y, err := strconv.ParseFloat(item.Y, 32)
		if err != nil {
			fmt.Println("Error : ", err)
			continue	
		}

		for _, rental := range item.Rental {
			
			numOfBedRoom, err := strconv.Atoi(rental.NoOfBedRoom)
			if err != nil {
				fmt.Println("Error : ", err)
				continue	
			}

			log.Println(rental.PropertyType)

			areaRange := strings.Split(rental.AreaSqm, "-")

			area, err := strconv.Atoi(areaRange[0])
			
			if err != nil {
				fmt.Println("Error : ", err)
				continue
			}
			foundProperty, err_query := dbclient.PropertyListing.
				Query().
				Where(
					propertylisting.And(
						propertylisting.NameEQ(item.Project),
						propertylisting.DistrictEQ(rental.District),
						propertylisting.HouseTypeEQ(rental.PropertyType),
						propertylisting.StreetEQ(item.Street),
						propertylisting.NumOfBedroom(numOfBedRoom),
					),
				).Limit(1).All(context.Background())

			if err_query == nil {
				p := foundProperty[0]
				price := p.RentalPrice
				n := p.NumberOfData
				p, err = p.Update().
					SetRentalPrice( (price * n + rental.Rent) / (n+1) ).
					SetNumberOfData(n+1).
					Save(context.Background())
				continue
			}

			u, err := dbclient.PropertyListing.
				Create().
					SetName(item.Project).
					SetHouseType(rental.PropertyType).
					SetArea(area).
					SetNumOfBedroom(numOfBedRoom).
					SetLeaseDate(rental.LeaseDate).
					SetRentalPrice(rental.Rent).
					SetX(x).
					SetY(y).
					SetStreet(item.Street).
					SetDistrict(rental.District).
					SetNumberOfData(1).
					Save(context.Background())

			if err != nil {
				log.Fatalf("[UpdatePropertyListing] : Error while updating db -> %s\n", err)
			}
			log.Println(u)
			log.Println("created")
		}
	}
}
