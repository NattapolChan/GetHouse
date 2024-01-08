import React, { useEffect, useState } from "react";
import Main from "../components/Main";
import ListView from "../components/ListView";
import MapView from "../components/MapView";
import { useNavigate } from "react-router-dom";
import SearchCriteria from "../components/SearchCriteria";

const DisplayInfoUI = ({ viewValue, toggleView }) => {
  const [navbar, setNavbar] = useState("NavbarDisplay");
  const navigate = useNavigate();

  const [marker, setMarkers] = useState([]);
  const [matchingItem, setMatchingItem] = useState([]);

  const handleDataFilteredMap = (filteredData1) => {
    setMarkers(filteredData1);
  };

  const handleDataFilteredList = (filteredData2) => {
    setMatchingItem(filteredData2);
  };

  useEffect(() => {
    // Check if we are on the /display route
    if (window.location.pathname === "/login") {
      // Set the navbar to 'NavbarDisplay' if on the /display route
      setNavbar("NavbarLogIn");
    } else if (window.location.pathname === "/") {
      setNavbar("NavbarHome");
    } else if (window.location.pathname === "/display") {
      setNavbar("NavbarDisplay");
    } else if (window.location.pathname === "/account") {
      setNavbar("NavbarAccount");
    } else if (window.location.pathname === "/signup") {
      setNavbar("NavbarSignUp");
    }
  }, []);

  useEffect(() => {
    // Check if we should navigate to /display
    if (navbar === "NavbarHome") {
      navigate("/");
    } else if (navbar === "NavbarDislay") {
      navigate("/display");
    } else if (navbar === "NavbarAccount") {
      navigate("/account");
    } else if (navbar === "NavbarLogIn") {
      navigate("/login");
    } else if (navbar === "NavbarSignUp") {
      navigate("/signup");
    }
  }, [navbar, navigate]);

  return (
    <>
      <Main />

      <SearchCriteria
        onDataFiltered1={handleDataFilteredMap}
        onDataFiltered2={handleDataFilteredList}
      />

      <div className="bg-base-100 font-mono">
        <div className="flex flex-col mx-4">
          {viewValue === "map" ? (
            <MapView markers={marker} matchingItems={matchingItem} />
          ) : (
            <ListView matchingItems={matchingItem} onDataFiltered2={handleDataFilteredList}/> 
          )}
        </div>
      </div>
    </>
  );
};

export default DisplayInfoUI;
