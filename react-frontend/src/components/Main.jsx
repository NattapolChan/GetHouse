import React, { useState, useEffect } from "react";

const Main = () => {
  const [bgImage, setBgImage] = useState("");

  // Define a function to get a random image URL
  const getRandomImage = () => {
    const imageNames = ["bg-1.jpg", "bg-2.jpg", "bg-3.jpg", "bg-4.jpg"];
    const randomIndex = Math.floor(Math.random() * imageNames.length);
    const randomImage = imageNames[randomIndex];
    const imageUrl = process.env.PUBLIC_URL + "/" + randomImage;
    return imageUrl;
  };

  useEffect(() => {
    // Set the background image when the component mounts
    setBgImage(getRandomImage());
  }, []);
return (<></>)

}

export default Main;
