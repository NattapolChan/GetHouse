import React, { useEffect, useState } from "react";
import House from "./House";

function FavoritesListing({ matchingItems }) {
  const [isLoaded, setIsLoaded] = useState(false);

  useEffect(() => {
    if (matchingItems.length) {
      setIsLoaded(true);
    } else {
      setIsLoaded(false);
    }
  }, [matchingItems]);
	console.log("favorite listing")

  return (
    <div className="flex flex-col mx-auto my-auto h-[800px] w-full justify-center ">
      {!isLoaded ? (
        <ul className="flex overflow-y-auto space-y-8 bg-gray-200 z-[91] mt-[-900px] ">
          <p>Loading...</p>
        </ul>
      ) : (
        <ul className="flex flex-col overflow-y-auto space-y-8 bg-gray-200 z-[91] rounded-[20px]  ">
          {matchingItems.map((property, index) => (
            <House viewValue={"list"} key={index} item={property} />
          ))}
        </ul>
      )}
    </div>
  );
}

export default FavoritesListing;
