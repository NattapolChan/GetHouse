import React, { useEffect, useState } from "react";
import {
  GoogleMap,
  useJsApiLoader,
  Marker,
  InfoWindow,
} from "@react-google-maps/api"; import { FaMapMarkerAlt, FaMapMarker } from "react-icons/fa";
import CustomMarker from "../assets/images/icons8-markerpurple.png";
import House from "./House";
import Cookies from "js-cookie";
import axios from "axios";

const containerStyle = {
  width: "900px",
  height: "450px",
};

const singaporeBounds = {
  // Define the LatLngBounds that encapsulate the entire Singapore area.
  north: 1.470771, // Adjust these coordinates as needed
  south: 1.202882,
  east: 104.05365,
  west: 103.607494,
};

function MapView({ markers, matchingItems }) {
  const [map, setMap] = useState(null);
  const [selectedMarker, setSelectedMarker] = useState(null);
  const [selectedIndex, setSelectedIndex] = useState(0);
  const [infoWindowOpen, setInfoWindowOpen] = useState(false);
  const [favList, setFavList] = useState([])
	console.log(process.env.REACT_APP_GOOGLE_MAP_API)
  const { isLoaded } = useJsApiLoader({
    id: "google-map-script",
    googleMapsApiKey: process.env.REACT_FIREBASE_APP_GOOGLE_MAP_API,
  });

  const onLoad = React.useCallback(function callback(map) {
    setMap(map);
  }, []);

  const onUnmount = React.useCallback(function callback(map) {
    setMap(null);
  }, []);

  useEffect(() => {
    if (isLoaded && map) {
      // Set the initial zoom level to fit the entire Singapore area.
      const bounds = new window.google.maps.LatLngBounds(singaporeBounds);
      map.fitBounds(bounds);
    }
  }, [isLoaded, map]);

  const handleMarkerMouseEnter = (marker, index) => {
    // Set the selected marker when mouse enters
    setSelectedMarker(marker);
    setSelectedIndex(index);
    setInfoWindowOpen(true);
  };

  const handleMarkerMouseLeave = () => {
    // Clear the selected marker when mouse leaves
    setSelectedMarker(null);
    setSelectedIndex("");
    setInfoWindowOpen(false);
  };

	const [token, setToken] = useState("");

	useEffect(() => {
		const cookiesToken = Cookies.get("token")

		if (cookiesToken) {
			setToken(cookiesToken)
		} else {
		}
	}, [])	

	const updateIdSet = (newId) => {
		setFavList(prevIds => new Set([...prevIds, newId]))
	}

	useEffect(() => {
		if (!token) return
		console.log(token)
		axios.get(`http://localhost:8080/get-favorite?accesstoken=${token}`)
			.then(res => {
				console.log(res.data) 
				res.data.forEach(item => updateIdSet(item.id))
			})
			.catch(err => {console.log(err)})
	}, [token])

  return  (
	  isLoaded ? (
    <div className='mx-auto h-full mt-4'>
      <GoogleMap
        mapContainerStyle={containerStyle}
        center={singaporeBounds}
        zoom={12}
        onLoad={onLoad}
        onUnmount={onUnmount}
        options={{
          mapTypeControl: false,
          streetViewControl: false,
          fullscreenControl: false,
        }}
      >
        {markers.map((marker, index) => (
          <Marker
            key={index}
            position={marker}
            options={{
              icon: CustomMarker,
            }}
            onClick={() => {handleMarkerMouseEnter(marker, index)}}
          ></Marker>
        ))}

        {selectedMarker && infoWindowOpen && matchingItems[selectedIndex] && (
          <InfoWindow
            position={selectedMarker}
            onCloseClick={() => {
              setSelectedMarker(null);
              setSelectedIndex("");
              setInfoWindowOpen(false);
              handleMarkerMouseLeave();
            }} // Handle closing the InfoWindow
          >
            <>
		{selectedIndex !== "" &&
		      <House 
			viewValue={"map"} 
			item={matchingItems[selectedIndex]}
			selectedIndex={selectedIndex}
			/>
		}
            </>
          </InfoWindow>
        )}
      </GoogleMap>
    </div>
  ) : (
    <></>)
  );
}

export default MapView;
