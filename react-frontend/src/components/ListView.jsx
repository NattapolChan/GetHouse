import React, { useState } from "react";
import House from "./House";
import {BiSolidDownArrow, BiSolidUpArrow} from "react-icons/bi"

function ListView({ matchingItems, onDataFiltered2 }) {

	const [activeSort, setActiveSort] = useState("");
	const [sortingOrder, setSortingOrder] = useState({
		column: "none", // Initially no column is selected for sorting
		ascending: true, // Initially sorting is in ascending order
	  });

	const handleColumnClick = (column) => {
		let newSortingOrder = { ...sortingOrder };
		if (!sortingOrder.ascending && sortingOrder.column===column) {
			setActiveSort("none")
			newSortingOrder.ascending = true
			setSortingOrder(newSortingOrder)
			return
		}
		if (column === newSortingOrder.column) {
		  newSortingOrder.ascending = !newSortingOrder.ascending;
		} else {
		  newSortingOrder.column = column;
		  newSortingOrder.ascending = true;
		}
		setSortingOrder(newSortingOrder);
		const sortedItems = sortItems(matchingItems, column, newSortingOrder.ascending);
		setActiveSort(column)
		onDataFiltered2(sortedItems)
	  };

	  const sortItems = (items, column, ascending) => {
		return items.slice().sort((a, b) => {
		  switch (column) {
			case "Project Name":
			  return ascending ? a.name.localeCompare(b.name) : b.name.localeCompare(a.name);
			case "Street":
			  return ascending ? a.street.localeCompare(b.street) : b.street.localeCompare(a.street);
			case "Number of Bedrooms":
			  return ascending ? a.num_of_bedroom - b.num_of_bedroom : b.num_of_bedroom - a.num_of_bedroom;
			case "Area (Square meter)":
				const areaComparison = parseFloat(a.area) - parseFloat(b.area);
				return ascending ? areaComparison : -areaComparison;
			case "Monthly Rental Price":
			  return ascending ? a.rental_price - b.rental_price : b.rental_price - a.rental_price;
			case "Property Type":
			  return ascending ? a.houseType.localeCompare(b.houseType) : b.houseType.localeCompare(a.houseType);
			default:
			  return 0;
		  }
		});
	  };

  return (
    <div className="overflow-x-auto mx-auto">
        <table className="table max-w-6xl">
	  <thead >
	  	<tr >
		  	<th ><span>Project Name</span>
	  			<button 
	  				className={activeSort==="Project Name" ? "mx-2 text-error" : "mx-2"} 
	  				onClick={() => handleColumnClick("Project Name")}>
	  			{sortingOrder.ascending && activeSort==="Project Name" ? <BiSolidDownArrow /> : <BiSolidUpArrow />}</button></th>
			<th><span>Street</span>
	  			<button 
	  				className={activeSort==="Street" ? "mx-2 text-error" : "mx-2"} 
	  				onClick={() => handleColumnClick("Street")}>
	  			{sortingOrder.ascending && activeSort==="Street" ? <BiSolidDownArrow /> : <BiSolidUpArrow />}</button></th>
			<th><span>Number of Bedrooms</span>
	  			<button 
	  				className={activeSort==="Number of Bedrooms" ? "mx-2 text-error" : "mx-2"} 
	  				onClick={() => handleColumnClick("Number of Bedrooms")}>
	  			{sortingOrder.ascending && activeSort==="Number of Bedrooms" ? <BiSolidDownArrow /> : <BiSolidUpArrow />}</button></th>
			<th><span>Area (Square meter)</span>
	  			<button 
	  				className={activeSort==="Area (Square meter)" ? "mx-2 text-error" : "mx-2"} 
	  				onClick={() => handleColumnClick("Area (Square meter)")}>
	  			{sortingOrder.ascending && activeSort==="Area (Square meter)" ? <BiSolidDownArrow /> : <BiSolidUpArrow />}</button></th>
			<th><span>Monthly Rental Price</span>
	  			<button 
	  				className={activeSort==="Monthly Rental Price" ? "mx-2 text-error" : "mx-2"} 
	  				onClick={() => handleColumnClick("Monthly Rental Price")}>
	  			{sortingOrder.ascending && activeSort==="Monthly Rental Price" ? <BiSolidDownArrow /> : <BiSolidUpArrow />}</button></th>
			<th><span>Property Type</span>
	  			<button 
	  				className={activeSort==="Property Type" ? "mx-2 text-error" : "mx-2"} 
	  				onClick={() => handleColumnClick("Property Type")}>
	  			{sortingOrder.ascending && activeSort==="Property Type" ? <BiSolidDownArrow /> : <BiSolidUpArrow />}</button></th>

            <th>Save to Fav</th>

	  	</tr>
	  </thead>
	  <tbody>
          {matchingItems.map((property, index) => (
            <House viewValue={"list"} key={index} selectedIndex={index} item={property} />
          ))}
	  </tbody>
        </table>
    </div>
  );
}

export default ListView;
