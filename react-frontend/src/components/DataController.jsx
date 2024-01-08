import React from "react";

const DataController = ({
  handleFindHouse,
  setSelectedLocation,
  setSelectedPriceRange,
  setSelectedNoOfRooms,
  setSelectedSize,
  setSelectedHouseType,
}) => {
  return (
<div className="bg-base-100 py-1">
      <div className="navbar space-x-4 mx-auto w-fit rounded-btn shadow-black">
	  <div className="navbar-center space-x-2 mx-auto">
		<div className="grid grid-col">
		  Location
		  <input type="text" placeholder="Type your location here" 
			className="input w-full max-w-xs input-accent input-sm" 
			onChange={(e) => {
				setSelectedLocation(e.target.value)
			}}/>
		</div>
		<div className="grid grid-col">
		  Price Range
		  <select className="select w-full max-w-xs select-accent select-sm" 
			onChange={(e) => {
				console.log(e.target.value)
				if (e.target.value===-1) {
					setSelectedPriceRange("")
					return
				}
				const priceLow = e.target.value * 500 + 1
				const priceHigh = e.target.value * 500 + 500
				setSelectedPriceRange(`${priceLow}-${priceHigh}`)
			}}>
			<option disabled selected>Price Range</option>
			<option value={-1}>NA</option>
			{Array.from(Array(20).keys()).map((num, _) => {
				return <option value={num}>{num*500+1}-{num*500+500}</option>
			})}
		  </select>
		</div>
		<div className="grid grid-col">
		  No. of Bedrooms
		  <select className="select max-w-xs select-accent select-sm" 
			onChange={(e) => {
				if (e.target.value === -1) setSelectedNoOfRooms(9999)
				else setSelectedNoOfRooms(e.target.value)
			}}>
			<option disabled selected>Number of Bedrooms</option>
			<option value={-1}>NA</option>
			{Array.from(Array(6).keys()).map((num, _) => {
				return <option value={num+1}>{num+1}</option>
			})}
		  </select>
		</div>
		<div className="grid grid-col">
		  Area
		  <input type="number" placeholder="Area (m^2)" 
			className="input w-full max-w-xs input-accent input-sm" 
			onChange={(e) => {
				setSelectedSize(e.target.value)
			}}/>
		</div>
		<div className="grid grid-col">
		  Property Type
		  <select className="select w-full max-w-xs select-accent select-sm" 
			onChange={(e) => {
				setSelectedHouseType(e.target.value)
			}}
		  >
			<option disabled selected>Property Type</option>
			<option value="Non-landed Properties">Non-landed Properties</option>
			<option value="Semi-Detached House">Semi-Detached House</option>
			<option value="Detached House">Detached House</option>
			<option value="Terrace House">Terrace House</option>
			<option value="Executive Condominium">Executive Condominium</option>
		  </select>
		</div>
		<button
		  onClick={handleFindHouse}
		  className="btn btn-accent mt-6 ml-0 btn-sm"
		>
		  Find Property
		</button>
	      </div>
	  </div>
	</div>
	  );
	};

export default DataController;
