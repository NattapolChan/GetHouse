import React, { useEffect } from "react";
import Cookies from "js-cookie";

const Redirect = () => {
	useEffect(() => {
		const urlParam = new URLSearchParams(window.location.search)
		const token = urlParam.get('token')
		if (token) {
			Cookies.set("token", token, {expires: 1});
		}

		window.location.href = '/account'
	}, [])

	return (<div className="bg-blue-300 text-white">Redirecting...</div>)
}

export default Redirect;
