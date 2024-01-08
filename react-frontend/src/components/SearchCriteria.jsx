import React, { useState } from "react";
import DataController from "./DataController";
import axios from "axios";
import { getDistance } from "geolib";
import svy21Wrapper from "./svy21Wrapper";

const SearchCriteria = ({ onDataFiltered1, onDataFiltered2 }) => {
  const [selectedLocation, setSelectedLocation] = useState("");
  const [selectedPriceRange, setSelectedPriceRange] = useState("");
  const [selectedNoOfRooms, setSelectedNoOfRooms] = useState("");

  const [selectedSize, setSelectedSize] = useState("");
  const [selectedHouseType, setSelectedHouseType] = useState("");

  const handleFindHouse = async () => {
    // Handle the logic when the "FindHouse" button is clicked
    const req = {
      location: selectedLocation,
      priceRange: selectedPriceRange,
      numberOfRooms: selectedNoOfRooms==="-1" ? "" : selectedNoOfRooms,
      size: selectedSize,
      houseType: selectedHouseType,
    };
	  console.log(req)

    // Construct the URL based on user criteria
	const apiUrl = `http://localhost:8080/search?location=${req.location}&priceRange=${req.priceRange}&houseType=${req.houseType}&size=${req.size}&numberOfRooms=${req.numberOfRooms}`;

    try {
      const response1 = await fetch(apiUrl);
      const result1 = await response1.json();
      console.log(apiUrl);
      console.log(result1);
      const matchingItemsArray = [];
      const markersArray = [];

      result1.forEach((item) => {
        //updateListMap.forEach((item)=>{})

        // Push the marker coordinates into the array
        const wgs84Coords = svy21Wrapper.computeLatLon(
          parseFloat(item.y),
          parseFloat(item.x),
        );
        markersArray.push({
          lat: wgs84Coords.lat,
          lng: wgs84Coords.lon,
        });

        // Store the matching rental items

        matchingItemsArray.push(item);
      });

      // Set the markers on the map
      onDataFiltered1(markersArray);
      onDataFiltered2(matchingItemsArray);
    } catch (error) {
      console.error("Error fetching data:", error);
    }

    // Make a fetch request to the constructed URL
  };

  return (
    <DataController
      handleFindHouse={handleFindHouse}
      setSelectedLocation={setSelectedLocation}
      setSelectedPriceRange={setSelectedPriceRange}
      setSelectedNoOfRooms={setSelectedNoOfRooms}
      setSelectedSize={setSelectedSize}
      setSelectedHouseType={setSelectedHouseType}
    />
  );
};

export default SearchCriteria;
