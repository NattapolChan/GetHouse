import React, { useEffect, useState } from "react";
import { Routes, Route, useNavigate } from "react-router-dom";
import Home from "./pages/Home";

import Login from "./pages/Login";
import Signup from "./pages/Signup";
import {
  GoogleMap,
  useLoadScript,
  useJsApiLoader,
  Marker,
} from "@react-google-maps/api";
import Redirect from "./pages/Redirect"
import DisplayInfoUI from "./pages/DisplayInfoUI";
import Account from "./pages/Account";
import BaseUIClass from "./components/BaseUICLass";
import { AuthContextProvider } from "./context/AuthContext";

function App() {
  const [view, setView] = useState("map");
  const [navbar, setNavbar] = useState("NavbarHome");
  const [textFromFile, setTextFromFile] = useState("");
  const [updateListMap, setUpdateListMap] = useState({});
  const [loggedIn, setLoggedIn] = useState(false);
  const navigate = useNavigate();

  const toggleView = async (selectedView) => {
    setView(selectedView);
  };
  const toggleNavbar = async (homeON) => {
    setNavbar(homeON);
  };
  const toggleLoggedIn = async (val) => {
    setLoggedIn(val);
  };
  useEffect(() => {
    // Check if we are on the /login route

    if (window.location.pathname === "/login") {
      // Set the navbar to 'NavbarLogIn' if on the /login route
      setNavbar("NavbarLogIn");
    } else if (window.location.pathname === "/") {
      setNavbar("NavbarHome");
      navigate("/");
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
    if (navbar === "NavbarDisplay") {
      navigate("/display");
    }
  }, [navbar, navigate]);

  return (
    <div className="">
      <AuthContextProvider>
        <BaseUIClass
          toggleView={toggleView}
          view={view}
          toggleNavbar={toggleNavbar}
          navbar={navbar}
          loggedIn={loggedIn}
          toggleLoggedIn={toggleLoggedIn}
        />

        <Routes>
          <Route path="/" element={<Home />} />
          <Route
            path="/display"
            element={<DisplayInfoUI viewValue={view} toggleView={toggleView} />}
          />
          <Route
            path="/login"
            element={
              <Login
                viewValue={view}
                toggleNavbar={toggleNavbar}
                loggedIn={loggedIn}
                toggleLoggedIn={toggleLoggedIn}
              />
            }
          />
          <Route
            path="/signup"
            element={
              <Signup
                viewValue={view}
                toggleNavbar={toggleNavbar}
                loggedIn={loggedIn}
                toggleLoggedIn={toggleLoggedIn}
              />
            }
          />
          <Route
            path="/account"
            element={<Account viewValue={view} toggleNavbar={toggleNavbar} />}
          />
	  <Route 
	    path="/redirect" 
	    element={<Redirect />}
	  />
        </Routes>
      </AuthContextProvider>
    </div>
  );
}

export default App;
