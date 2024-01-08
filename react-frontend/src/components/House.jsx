import React, { useState, useEffect } from "react";
import { GoBookmark, GoBookmarkFill } from "react-icons/go";
import Cookies from "js-cookie";
import axios from "axios";

const House = (props) => {
	const item = props.item
	const viewValue = props.viewValue
	const index = props.selectedIndex
	const [favList, setFavList] = useState(new Set())

	const [token, setToken] = useState("")

	useEffect(() => {
		setToken(Cookies.get("token"))
	}, [])

	useEffect(() => {
		if (!token) return
		axios.get(`http://localhost:8080/get-favorite?accesstoken=${token}`)
			.then(res => {
				setFavList(new Set())
				res.data?.forEach(item => updateIdSet(item.id))
			})
			.catch(err => console.log(err))
	}, [token, index])

	const removeFavorite = () => {
		if (!token) return
		if (!item.id) return
		axios.get(`http://localhost:8080/remove-favorite?accesstoken=${token}&listid=${item.id}`)
			.then(res => {
				setFavList(new Set())
				res.data?.forEach(item => updateIdSet(item.id))
			})
			.catch(err => console.log(err))
	}

	const addFavorite = () => {
		if (!token) return
		if (!item.id) return
		axios.get(`http://localhost:8080/add-favorite?accesstoken=${token}&listid=${item.id}`)
			.then(res => res.data?.forEach(item => updateIdSet(item.id)))
			.catch(err => console.log(err))
	}

	const updateIdSet = (newId) => {
		setFavList(prevIds => new Set([...prevIds, newId]))
	}

	useEffect(() => {
		if (!token) return
		axios.get(`http://localhost:8080/get-favorite?accesstoken=${token}`)
			.then(res => {
				res.data.forEach(item => updateIdSet(item.id))
			})
			.catch(err => {console.log(err)})
	}, [token])

	useEffect(() => {
	}, [favList, item])

  return viewValue === "list" ? ( 
	<tr>
		<th>{item?.name}</th>
	  	<th>{item?.street}</th>
	  	<th>{item?.num_of_bedroom}</th>
	  	<th>{item?.area}</th>
	  	<th>{item?.rental_price}</th>
	  	<th>{item?.houseType}</th>
	  	<th>{ Array.from(favList).some(i => i===item.id) ? (
			<button className="text-secondary" onClick={() => {removeFavorite()}}>
				<GoBookmarkFill size={28} />
			</button> ) :
			(
			<button onClick={() => {addFavorite()}}>
				<GoBookmark size={28} />
			</button>
			)
		}</th>
	</tr>
  ) : (
    <>
      <div className="stats bg-white m-2 text-gray-700 font-bold font-sans">
	<ul>
          <li className="badge text-white text-xl font-extrabold p-4 rounded-btn">
            {item?.name}
          </li>
	  <li className="p-1">Street: {item?.street}</li>
          <li className="p-1">House type: {item?.houseType}</li>
          <li className="p-1">Monthly Rental Price: {item?.rental_price}</li>
          <li className="p-1">Number of bedrooms: {item?.num_of_bedroom}</li>
          <li className="p-1">
            Size {"("}areaSqm{")"}: {item?.area}
	  </li>
	</ul>
        <div>
          {Array.from(favList).some(i => i===item.id) ? (
            // Display GoBookmarkFill if the property is in favorites
            <button className="text-secondary" onClick={() => {removeFavorite()}}>
              <GoBookmarkFill size={28} />
            </button>
          ) : (
            // Display GoBookmark if the property is not in favorites
            <button onClick={() => {addFavorite()}}>
              <GoBookmark size={28} />
            </button>
          )}
        </div>
      </div>
    </>
  );
};
export default House;
