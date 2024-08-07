// APIs like ipinfo.io, ipstack, or ipgeolocation.io.
// src/components/GeoLocation.tsx (updated)
"use client";

import React, { useEffect, useState } from 'react';

const GeoLocation = () => {
  const [location, setLocation] = useState(null);

  useEffect(() => {
    // Replace 'YOUR_API_KEY' with the actual key for the geolocation service you are using
    fetch('https://ipinfo.io/json?token=YOUR_API_KEY')
      .then((response) => response.json())
      .then((data) => setLocation(data))
      .catch((error) => console.error('Error fetching geolocation:', error));
  }, []);

  if (!location) {
    return <div>Loading...</div>;
  }

  return (
    <div>
      <h3>Location Information</h3>
      <p>IP: {location.ip}</p>
      <p>Country: {location.country}</p>
      <p>Region: {location.region}</p>
      <p>City: {location.city}</p>
    </div>
  );
};

export default GeoLocation;
