import React, { useEffect, useState } from "react";
import Cookies from "js-cookie";

const BaseUIClass = ({
  toggleView,
  toggleNavbar,
  view,
  navbar,
  loggedIn,
  toggleLoggedIn,
}) => {
  const [views, setViews] = useState(view);

	const [token, setToken] = useState("");

	useEffect(() => {
		const cookiesToken = Cookies.get("token")

		if (cookiesToken) {
			setToken(cookiesToken)
		} else {
			//window.location.href = '/login'
		}
	}, [])

  const handleLogout = () => {
	Cookies.remove("token")
	setToken("")
	setAlertVisible(true)
  };

	const [isAlertVisible, setAlertVisible] = useState(false)

  return (
    <>
        <div className="z-50 font-mono">
		{navbar !== "NavbarDisplay" ? (
			<div className="navbar">
				<div className="navbar-start m-3 space-x-3">
					<a 
					className="btn btn-sm btn-primary text-sm" 
					href="/account"
					onClick={() => {
						toggleNavbar("NavbarAccount")
					}}
					>{token ? "Favorites" : "login"}</a>
					{ token && <button
					className="btn btn-sm btn-error text-sm"
					onClick={handleLogout}
					>Logout</button> }
				</div>
				<div className="navbar-center text-xl font-bold">
					<a
					className="btn-xl"
					onClick={()=> {
					    toggleNavbar("NavbarHome")
					}}
					href="/"
					>
					Welcome to GetHouse
					</a>
				</div>
				<div className="navbar-end">
					<button className="btn btn-ghost btn-circle"
						onClick={() => {
						      toggleView("map"); // Call toggleView with 'map' to switch to Map component
						      toggleNavbar("NavbarDisplay");
						}}>
						<svg xmlns="http://www.w3.org/2000/svg" 
						className="h-5 w-5" fill="none" 
						viewBox="0 0 24 24" stroke="currentColor"
						><path strokeLinecap="round" 
						strokeLinejoin="round" strokeWidth="2" 
						d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
					</button>
				</div>
			</div>
		): (
			<div className="navbar">
				<div className="navbar-start m-1 space-x-3">
					<a 
					className="btn btn-sm btn-primary text-sm" 
					href="/account"
					onClick={() => {
						toggleNavbar("NavbarAccount")
					}}
					>{token ? "Favorites" : "login"}</a>
					{ token && <button
						className="btn btn-sm btn-error text-sm"
						onClick={handleLogout}
					>Logout</button>}
				</div>
				<div className="navbar-center text-xl font-bold">
					<button
					className="btn-xl"
					onClick={()=> {
					    toggleNavbar("NavbarHome")
					}}
					>
					Welcome to GetHouse
					</button>
				</div>
				<div className="navbar-end">
					<button className="btn btn-ghost btn-circle"
						onClick={() => {
						      toggleView("list"); // Call toggleView with 'map' to switch to Map component
						      toggleNavbar("NavbarDisplay");
						}}>
						<svg xmlns="http://www.w3.org/2000/svg" 
						className="h-5 w-5" fill="none" 
						viewBox="0 0 24 24" stroke="currentColor"
						><path strokeLinecap="round" 
						strokeLinejoin="round" strokeWidth="2" 
						d="M21 21l-6-6m2-5a7 7 0 11-14 0 7 7 0 0114 0z" /></svg>
					</button>
					<div className="tabs tabs-boxed px-4">
						<button
							className={views==="map" ? "tab" : "tab tab-active"}
							onClick={() => {setViews("list"); toggleView("list")}}
						>List</button> 
						<button 
							className={views==="map" ? "tab tab-active" : "tab"}
							onClick={() => {setViews("map"); toggleView("map")}}
						>Map</button>
					</div>
				</div>
			</div>
			)
		}
        </div>
	{isAlertVisible && (
		<div className="alert alert-success p-1">
		  <svg xmlns="http://www.w3.org/2000/svg" className="stroke-current shrink-0 h-6 w-6" 
			fill="none" viewBox="0 0 24 24">
			<path strokeLinecap="round" 
			strokeLinejoin="round" strokeWidth="2" 
			d="M9 12l2 2 4-4m6 2a9 9 0 11-18 0 9 9 0 0118 0z" /></svg>
		  <span>You have been logged out!</span>
		  <button className="btn btn-ghost" onClick={()=> {setAlertVisible(false)}}>Close</button>
		</div>
      )}
    </>
  );
};

export default BaseUIClass;
