import React, { useEffect, useState } from "react";
import Main from "../components/Main";
import Icon from "../components/Icon";
import { useNavigate } from "react-router-dom";
import BaseUIClass from "../components/BaseUICLass";
import MapView from "../components/MapView";
import { UserAuth } from "../context/AuthContext";
import { doc, onSnapshot, updateDoc } from "firebase/firestore";
import { db } from "../firebase";
import ListView from "../components/ListView";
import svy21Wrapper from "../components/svy21Wrapper";
import Cookies from "js-cookie";
import axios from "axios";

const Account = ({ viewValue, toggleView }) => {
  const [propertys, setPropertys] = useState([]);
  const { user } = UserAuth();
  const [marker1, setMarker1] = useState([]);
  const [matchingItems1, setMatchingItems1] = useState([]);

  const [navbar, setNavbar] = useState("NavbarAccount");
  const navigate = useNavigate();

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

/*
  useEffect(() => {
    onSnapshot(doc(db, "users", `${user?.email}`), (doc) => {
      setPropertys(doc.data()?.savedProperty);
    });
  }, [user?.email]);
	*/

	const [token, setToken] = useState("");

	useEffect(() => {
		const cookiesToken = Cookies.get("token")

		if (cookiesToken) {
			setToken(cookiesToken)
		} else {
			window.location.href = '/login'
		}
	}, [])	

	useEffect(() => {
		if (!token) return
		console.log(token)
		axios.get(`http://localhost:8080/get-favorite?accesstoken=${token}`)
			.then(res => {console.log(res.data); setPropertys(res.data)})
			.catch(err => {console.log(err)})
	}, [token])

  useEffect(() => {
    const markersArray = [];
    const matchingItemsArray = [];

	console.log("int prop")

    propertys?.forEach((item) => {
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
      matchingItemsArray.push(propertys);
    });
    setMarker1(markersArray);
    setMatchingItems1(matchingItemsArray);
	  console.log(matchingItemsArray)
  }, [propertys]);

  return (
    <>
      <Main />
      <div className="w-full h-screen mx-auto text-center flex flex-col justify-center">
	  {matchingItems1 && <MapView markers={marker1} matchingItems={matchingItems1[0]} />}
      </div>
    </>
  );
};

export default Account;
